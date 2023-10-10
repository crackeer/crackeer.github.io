package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	list := os.Args
	if len(list) < 2 {
		return
	}
	download(list[1])
}

func download(urlString string) error {
	object, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	target := filepath.Join("./result", object.Path)
	if value := os.Getenv("ROOT_PATH"); len(value) > 0 {
		target = filepath.Join(value, object.Path)
	}

	return downloadTo(urlString, target)

}

func downloadTo(urlString string, target string) error {
	fmt.Println("Downloading-->", urlString, target)
	dir, _ := filepath.Split(target)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	response, err := http.Get(urlString)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(target, bytes, os.ModePerm)
}
