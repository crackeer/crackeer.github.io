package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
var protocol string = `[A-Z]+://`
var hostname string = `[-\\w]+(?:\\.\\w[-\\w]*)+`
var port string = `:\\d+`
var path string = `/[^.!,?\"<>\\[\\]{}\\s\\x7F-\\xFF]*` + `(?:[.!,?]+[^.!,?\"<>\\[\\]{}\\s\\x7F-\\xFF]+)*`
*/

func main() {
	bytes, _ := os.ReadFile(os.Args[1])

	var dest interface{}
	if err := json.Unmarshal(bytes, &dest); err != nil {
		panic(err.Error())
	}
	list := extractURL(dest)
	fmt.Println(list)

}
func isURL(value string) bool {
	return strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://")
}

func extractURL(value interface{}) []string {
	if strValue, ok := value.(string); ok && isURL(strValue) {
		return []string{strValue}
	}

	var retData []string
	if mapValue, ok := value.(map[string]interface{}); ok {
		for _, value := range mapValue {
			retData = append(retData, extractURL(value)...)
		}
		return retData
	}

	if listValue, ok := value.([]interface{}); ok {
		for _, value := range listValue {
			retData = append(retData, extractURL(value)...)
		}
		return retData
	}
	return nil
}
