package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Args[3])

	var (
		url    string = os.Args[2]
		method string = os.Args[1]
		header map[string]string
		file   string = os.Args[3]
		err    error
		bodys  []string
	)

	header, err = readHeader()
	if err != nil {
		panic(err.Error())
	}
	bodys, err = readFile(file)
	if err != nil {
		panic(err.Error())
	}

	doHttp(method, url, bodys, header)
}

func doHttp(method, url string, body []string, header map[string]string) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	for _, item := range body {
		reader := strings.NewReader(item)
		request, err := http.NewRequest(method, url, reader)
		if method == http.MethodGet {
			request, err = http.NewRequest(method, url+"?"+item, nil)
		}
		if err != nil {
			panic(err.Error())
		}
		response, err := client.Do(request)
		if err != nil {
			panic(err.Error())
		}
		bytes, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(bytes))
	}
}

func readHeader() (map[string]string, error) {
	f, exist := os.LookupEnv("HEADER_FILE")
	if !exist {
		return nil, nil
	}
	bytes, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	retData := map[string]string{}
	err = json.Unmarshal(bytes, &retData)
	return retData, err

}

func readFile(file string) ([]string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("open file %s error:%s", file, err.Error())
	}

	return strings.Split(string(bytes), "\n"), nil
}
