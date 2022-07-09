# 遍历文件夹

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	list := getAllFile("d:/github/go-code-test")
	for _, v := range list {
		fmt.Println(v)
	}
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
```