package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var jsonX = jsoniter.ConfigFastest

func main() {
	a := map[string]interface{}{
		"abc": "&#*?|!@#$%^&*()'\"",
	}
	bytes, _ := jsonX.Marshal(a)
	fmt.Println(string(bytes))
}
