package main

import (
	"fmt"
	"strconv"
)

func main() {

	from := 1 << 7
	to := 1 << 5

	fmt.Println(from&to > 0, from&to)
	from |= 1 << 5
	from |= 1 << 3
	fmt.Println(strconv.FormatInt(int64(from), 2))

	aa := ^(1 << 5)
	fmt.Println(strconv.FormatInt(int64(1<<5), 2))
	fmt.Println(aa)
	fmt.Println(strconv.FormatInt(int64(aa), 2))
	from &= ^(1 << 5)
	fmt.Println(strconv.FormatInt(int64(from), 2))
}
