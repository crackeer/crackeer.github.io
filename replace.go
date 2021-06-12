package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	//"io"
	"os"
	//"sync"
)

var downloadDir = "./upload/iloli"

func checkDownloadDir() {
	_, err := os.Stat(downloadDir)

	if os.IsNotExist(err) {
		os.MkdirAll(downloadDir, 0777)
		os.Chmod(downloadDir, 0777)
	}
}

func main() {
	checkDownloadDir()
	l := getMDList(`./story`)
	for _, f := range l {
		parseImage(f)
	}
}

func parseImage(f string) {
	content, _ := ioutil.ReadFile(f)
	eg, err := regexp.Compile(`https:\/\/i.loli.net\/\d{4}\/\d{2}/\d{2}\/\S+.png`)
	if err != nil {
		panic(err)
	}
	newContent := eg.ReplaceAllFunc(content, func(dest []byte) []byte {
		parts := strings.Split(string(dest), "/")

		return []byte(fmt.Sprintf("https://gitee.com/seekgo/image/raw/master/%s", parts[len(parts)-1]))
	})

	fmt.Println(string(newContent))
	ioutil.WriteFile(f, newContent, os.ModePerm)
}

func getMDList(folder string) []string {

	files, _ := ioutil.ReadDir(folder)
	retData := []string{}
	for _, file := range files {
		if file.IsDir() {
			tmp := getMDList(fmt.Sprintf("%s/%s", folder, file.Name()))
			retData = append(retData, tmp...)
		} else {
			if strings.HasSuffix(file.Name(), ".md") {
				retData = append(retData, fmt.Sprintf("%s/%s", folder, file.Name()))
			}

		}
	}
	return retData
}
