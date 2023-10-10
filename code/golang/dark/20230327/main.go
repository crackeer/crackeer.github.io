package main

import "fmt"

func main() {
	list := catData([]byte("abc"))
	fmt.Println(list)
}

func catData(bytes []byte) []string {
	if len(bytes) < 1 {
		return []string{}
	}

	if len(bytes) < 2 {
		return []string{string(bytes)}
	}

	result := catData(bytes[1:])
	retData := []string{}
	retData = append(retData, result...)
	for _, item := range result {
		tmp := make([]byte, len(item)+1)
		tmp[0] = bytes[0]
		copy(tmp[1:], []byte(item))
		retData = append(retData, string(tmp))
	}
	retData = append(retData, string(bytes[0:1]))
	return retData
}
