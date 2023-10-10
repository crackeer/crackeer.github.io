package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.OpenFile("/tmp/goalang_file_append.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	for i := 0; i < 1000; i++ {
		file.WriteString(fmt.Sprintf("%d\n", i*18))
	}

	/*
		list := getAllFile("d:/github/go-code-test")
		for _, v := range list {
			fmt.Println(v)
		}
	*/
}

func getAllFile(folder string) []string {
	files, _ := ioutil.ReadDir(folder)
	retData := []string{}
	for _, file := range files {
		if file.IsDir() {
			tmp := getAllFile(fmt.Sprintf("%s/%s", folder, file.Name()))
			retData = append(retData, tmp...)
		} else {
			retData = append(retData, fmt.Sprintf("%s/%s", folder, file.Name()))
		}
	}
	return retData
}
