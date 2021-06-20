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
	l := getMDList(`./diary`)
	for _, f := range l {
		replaceImage(f)
	}
}

func replaceImage(f string) {

	content, _ := ioutil.ReadFile(f)
	eg, err := regexp.Compile(`https:\/\/gitee.com\/seekgo\/image\/raw\/master\/\S+.png`)
	if err != nil {
		panic(err)
	}

	newContent := eg.ReplaceAllFunc(content, func(dest []byte) []byte {
		parts := strings.Split(string(dest), "/")
		return []byte(fmt.Sprintf("image/%s", parts[len(parts)-1]))

	})

	ioutil.WriteFile(f, newContent, os.ModePerm)
}

func findImages(f string) []string {
	parts := strings.Split(f, "/")
	dir := strings.Join(parts[0:len(parts)-1], "/")
	imageDir := fmt.Sprintf("%s/image", dir)

	content, _ := ioutil.ReadFile(f)
	eg, err := regexp.Compile(`https:\/\/gitee.com\/seekgo\/image\/raw\/master\/\S+.jpg`)
	if err != nil {
		panic(err)
	}

	images := eg.FindAllString(string(content), -1)

	if len(images) > 0 {
		DownloadImage(images, imageDir)
	}

	return images
}

func DownloadImage(images []string, dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	dest := "D:/gitee/image/"

	for _, v := range images {

		fmt.Println("Download", v)

		tmpImageName := strings.Replace(v, "https://gitee.com/seekgo/image/raw/master/", "", -1)
		destImagePath := fmt.Sprintf("%s/%s", dest, tmpImageName)

		parts := strings.Split(v, "/")
		imageName := parts[len(parts)-1]
		imageFullPath := fmt.Sprintf("%s/%s", dir, imageName)

		if _, err := os.Stat(destImagePath); os.IsNotExist(err) {
			fmt.Println("Not Found", v)
			continue
		}

		if _, err := os.Stat(imageFullPath); os.IsNotExist(err) {

			data, _ := ioutil.ReadFile(destImagePath)
			ioutil.WriteFile(imageFullPath, data, os.ModePerm)
		}
	}
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
