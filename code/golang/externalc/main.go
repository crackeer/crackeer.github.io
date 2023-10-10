package main

import (
	"fmt"
	"time"

	"github.com/crackeer/golang-code/externalc/safe"
)

func main() {
	guard, err := safe.NewDefaultSafeGuard()
	if err != nil {
		panic(err.Error())
		return
	}
	info, err := guard.GetLicenseInfo()
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println(string(info.JSON()))

	lockInfo, err := guard.GetLockInfo()
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println("LockINFO-->LOCKSN", lockInfo)

	fmt.Println("Write Raw Data-->")
	_, err = guard.WriteRawData(fmt.Sprintf("StartAt:%d", time.Now().Unix()))
	if err != nil {
		fmt.Println("WriteRawDataError", err.Error())
		return

	}

	rawText, err := guard.ReadRawData()
	if err != nil {
		fmt.Println("ReadRawDataError", err.Error())
		return
	}
	fmt.Println("RAW:", rawText)
	romText, err := guard.ReadRomData()
	if err != nil {
		fmt.Println("ReadRomDataError", err.Error())
		return
	}
	fmt.Println("ROM:", romText)

	pubText, err := guard.ReadPubData()
	if err != nil {
		fmt.Println("ReadPubDataError", err.Error())
		return
	}
	fmt.Println("PUB:", pubText)

	fmt.Println("Logout:", guard.Logout())

}
