package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	decoder := json.NewDecoder(bytes.NewReader([]byte("103113aa2e7dadaa3c7f439124f7e183844567-fail")))
	var someData interface{}
	err := decoder.Decode(&someData)
	fmt.Println(err == nil)

	err = json.Unmarshal([]byte("103113aa2e7dadaa3c7f439124f7e183844567-fail"), &someData)
	fmt.Println(err == nil)
}
