package main

import (
	"fmt"
	"regexp"
)

var userCodeRegex = regexp.MustCompile(`^[0-9a-zA-Z]{13}$`)

func main() {
	for _, item := range []string{
		"4W1aZNYkrOz3Vs", "67575", "4W1aZNYkrOz3-",
	} {
		flag := userCodeRegex.Match([]byte(item))
		fmt.Println(item, flag)
	}

	abc := "save to message error[code=1372]"

	exp, err := regexp.Compile(`\[code=(\d+)\]$`)
	fmt.Println(err)

	fmt.Println(exp.FindStringSubmatch(abc))

}
