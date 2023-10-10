package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	for _, item := range list[1:] {
		content := readFile(item)
		fmt.Println(content)
		fmt.Println("")
	}
}

func readFile(file string) string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}
