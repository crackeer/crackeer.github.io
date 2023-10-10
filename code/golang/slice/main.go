package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2}
	foo(a)
	fmt.Println(a)

	slice1 := make([]int, 0, 100)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := slice1
	slice1[1] = 99
	fmt.Println(slice1, slice2)
	fmt.Println(fmt.Sprintf("%p", slice1))
	fmt.Println(fmt.Sprintf("%p", slice2))
	testString()
}

func testString() {
	str := "abc"
	str1 := str
	str = "def"
	fmt.Println(str, str1)
}

func foo(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
}
