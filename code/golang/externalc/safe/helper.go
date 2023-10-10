package safe

import "C"

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func getUintValue(b byte) uint8 {
	mapByte := []byte("0123456789ABCDEF")
	for index, item := range mapByte {
		if item == b {
			return uint8(index)
		}
	}
	return 0
}

func convert2Uint(pair []byte) uint8 {
	var retData uint8
	for i := len(pair) - 1; i >= 0; i-- {
		retData = retData + getUintValue(pair[i])*uint8(math.Pow(16, float64(len(pair)-1-i)))
	}
	return retData
}

func Convert32To16(data string) ([]byte, error) {
	if len(data) != 32 {
		return []byte{}, errors.New("password length not 32")
	}
	// Create a [16]unsigned char array
	var array []byte = make([]byte, 16)
	for i := range array {
		tt := convert2Uint([]byte(data[i*2 : (i+1)*2]))
		array[i] = byte(tt)
	}
	fmt.Println(len(array))
	return array, nil
}

func convertPasswd(slmInitPasswd string) ([16]C.uchar, error) {

	if len(slmInitPasswd) != 32 {
		return [16]C.uchar{}, errors.New("password length not 32")
	}
	// Create a [16]unsigned char array
	var array [16]C.uchar
	for i := range array {
		tt := convert2Uint([]byte(slmInitPasswd[i*2 : (i+1)*2]))
		array[i] = C.uchar(tt)
	}
	return array, nil
}

func toErrorCode(result C.uint) string {
	return fmt.Sprintf("0x%s", strconv.FormatInt(int64(result), 16))
}
