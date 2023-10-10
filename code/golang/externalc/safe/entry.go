package safe

// #cgo CFLAGS: -I/usr/local/share/senseshield/sdk/C/include
// #cgo LDFLAGS: -L/usr/local/share/senseshield/sdk/C/lib64 -lslm_runtime -lm -lpthread
// #include "ss_lm_runtime.h"
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"unsafe"
)

// SafeGuard
type SafeGuard struct {
	Password  string
	Timeout   uint
	LicenseID uint

	handleIndex C.uint
}

// LicenseInfo ...
type LicenseInfo struct {
	LicenseID           int    `json:"license_id"`
	Enable              bool   `json:"enable"`
	GUID                string `json:"guid"`
	StartTime           int    `json:"start_time"`
	EndTime             int    `json:"end_time"`
	FirstUseTime        int    `json:"first_use_time"`
	ConcurrentType      string `json:"concurrent_type"`
	Concurrent          int    `json:"concurrent"`
	Version             int    `json:"version"`
	Module              string `json:"module"`
	LastUpdateTimestamp int    `json:"last_update_timestamp"`
	LastUpdateTimeSN    int    `json:"last_update_timesn"`
	RomSize             int    `json:"rom_size"`
	RawSize             int    `json:"raw_size"`
	PubSize             int    `json:"pub_size"`
	DeveloperID         string `json:"developer_id"`
	Type                string `json:"type"`
	Sn                  string `json:"sn"`
}

type LockInfo struct {
	Clock                 int    `json:"clock"`
	AvailableSpace        int    `json:"available_space"`
	TotalSpace            int    `json:"total_space"`
	CommunicationProtocol int    `json:"communication_protocol"`
	LockFirmwareVersion   string `json:"lock_firmware_version"`
	LmFirmwareVersion     string `json:"lm_firmware_version"`
	H5DeviceType          int    `json:"h5_device_type"`
	ClockType             int    `json:"clock_type"`
	SharedEnabled         int    `json:"shared_enabled"`
	OwnerDeveloperID      string `json:"owner_developer_id"`
	DeviceModel           string `json:"device_model"`
	HardwareVersion       string `json:"hardware_version"`
	ManufactureDate       string `json:"manufacture_date"`
	LockSn                string `json:"lock_sn"`
	SlaveAddr             int    `json:"slave_addr"`
	ShellNum              string `json:"shell_num"`
	UserInfo              string `json:"user_info"`
	InnerInfo             string `json:"inner_info"`
}

// Available ...
//
//	@receiver info
//	@return bool
func (info *LicenseInfo) Available() bool {
	nowTs := int(time.Now().Unix())
	return nowTs >= info.StartTime && nowTs <= info.EndTime && info.Enable
}

// JSON
//
//	@receiver info
//	@return []byte
func (info *LicenseInfo) JSON() []byte {
	bytes, _ := json.Marshal(info)
	return bytes
}

// NewDefaultSafeGuard
//
//	@return *SafeGuard
//	@return error
func NewDefaultSafeGuard() (*SafeGuard, error) {
	guardPasswd, exists := os.LookupEnv("GUARD_PASSWD")
	if !exists {
		return nil, errors.New("env GUARD_PASSWD not exists")
	}
	guardLicenseID, exists := os.LookupEnv("GUARD_LICENSE_ID")
	if !exists {
		return nil, errors.New("env GUARD_LICENSE_ID not exists")
	}
	licenseID, _ := strconv.Atoi(guardLicenseID)

	guardTimeout, exists := os.LookupEnv("GUARD_TIMEOUT")
	if !exists {
		guardTimeout = "600"
	}

	timeout, _ := strconv.Atoi(guardTimeout)
	if timeout < 1 || timeout < 60 {
		timeout = 60
	}
	return NewSafeGuard(guardPasswd, uint(licenseID), uint(timeout))
}

// NewSafeGuard
//
//	@param passwd
//	@param licenseID
//	@param timeout
//	@return *SafeGuard
//	@return error
func NewSafeGuard(passwd string, licenseID uint, timeout uint) (*SafeGuard, error) {
	// ------ 1. slm_init ---------
	var initParams C.ST_INIT_PARAM
	initParams.version = C.SLM_CALLBACK_VERSION02
	initParams.pfn = nil
	if value, err := convertPasswd(passwd); err != nil {
		return nil, err
	} else {
		initParams.password = value
	}

	result := C.slm_init(&initParams)
	if result > 0 {
		return nil, fmt.Errorf("init slm error, error_code = %s", toErrorCode(result))
	}

	// ------ 2. slm_login -----
	var loginParams C.ST_LOGIN_PARAM
	loginStructSize := C.uint(unsafe.Sizeof(C.ST_LOGIN_PARAM{}))
	loginParams.license_id = C.uint(licenseID)
	loginParams.size = loginStructSize
	loginParams.timeout = C.uint(timeout)
	loginParams.login_mode = C.SLM_LOGIN_MODE_AUTO
	handleIndex := C.uint(0)
	result = C.slm_login(&loginParams, C.STRUCT, &handleIndex, nil)
	if result > 0 {
		return nil, fmt.Errorf("slm_login error, error_code = %s", toErrorCode(result))
	}

	return &SafeGuard{
		Password:    passwd,
		Timeout:     timeout,
		LicenseID:   licenseID,
		handleIndex: handleIndex,
	}, nil
}

// GetLicenseInfo
//
//	@receiver guard
//	@return *LicenseInfo
//	@return error
func (guard *SafeGuard) GetLicenseInfo() (*LicenseInfo, error) {
	var licenseInfo *C.char
	result := C.slm_get_info(guard.handleIndex, C.LICENSE_INFO, C.JSON, &licenseInfo)
	if result > 0 {
		return nil, fmt.Errorf("slm_get_info error, error_code = %s", toErrorCode(result))
	}
	retData := &LicenseInfo{}
	if err := json.Unmarshal([]byte(C.GoString(licenseInfo)), retData); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %s", err.Error())
	}
	return retData, nil
}

func (guard *SafeGuard) GetLockInfo() (*LockInfo, error) {
	var lockInfo *C.char
	result := C.slm_get_info(guard.handleIndex, C.LOCK_INFO, C.JSON, &lockInfo)
	if result > 0 {
		return nil, fmt.Errorf("slm_get_info error, error_code = %s", toErrorCode(result))
	}
	retData := &LockInfo{}
	if err := json.Unmarshal([]byte(C.GoString(lockInfo)), retData); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %s", err.Error())
	}
	return retData, nil
}

// ReadRomData
//
//	@receiver guard
//	@return string
//	@return error
func (guard *SafeGuard) ReadRomData() (string, error) {
	return guard.readData(C.ROM)
}

// ReadRawData ...
//
//	@receiver guard
//	@return string
//	@return error
func (guard *SafeGuard) ReadRawData() (string, error) {
	return guard.readData(C.RAW)
}

// ReadPubData
//
//	@receiver guard
//	@return string
//	@return error
func (guard *SafeGuard) ReadPubData() (string, error) {
	return guard.readData(C.PUB)
}

func (guard *SafeGuard) readData(dataType C.LIC_USER_DATA_TYPE) (string, error) {
	dataLen, err := guard.GetUserDataSize(dataType)
	if err != nil {
		return "", err
	}
	if dataLen < 1 {
		return "", nil
	}
	var rawData *C.uchar = new(C.uchar)
	result := C.slm_user_data_read(guard.handleIndex, dataType, rawData, 0, dataLen)
	if result > 0 {
		return "", fmt.Errorf("slm_user_data_read error_code: [%d], %s", result, toErrorCode(result))
	}
	data := C.GoString((*C.char)(unsafe.Pointer(rawData)))
	return string(data), nil
}

func (guard *SafeGuard) GetUserDataSize(dataType C.LIC_USER_DATA_TYPE) (C.uint, error) {
	var dataLen C.uint
	result := C.slm_user_data_getsize(guard.handleIndex, dataType, &dataLen)
	if result > 0 {
		return 0, fmt.Errorf("slm_user_data_getsize error_code [%d], %s", result, toErrorCode(result))
	}
	return dataLen, nil
}

// Logout
//
//	@receiver guard
//	@return uint
func (guard *SafeGuard) Logout() uint {
	return uint(C.slm_logout(guard.handleIndex))
}

func (guard *SafeGuard) WriteRawData(data string) (uint, error) {

	rawSize, err := guard.GetUserDataSize(C.RAW)
	if err != nil {
		return 1, fmt.Errorf("get user data size error:%s", err.Error())
	}

	dataLen := C.size_t(len([]byte(data)))
	if uint(rawSize) < uint(dataLen) {
		return 2, errors.New("no enough space to write the data")
	}

	result := C.slm_user_data_write(guard.handleIndex, cUCharStr, 0, C.uint(dataLen))
	if result != 0 {
		return uint(result), fmt.Errorf("write raw data error_code = %s", toErrorCode(result))
	}
	return 0, nil
}
