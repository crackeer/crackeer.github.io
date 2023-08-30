package main

import (
	"fmt"
	"math"
	"strings"
)

func foo(slice []int) {
	slice[0] = 100
}

func main() {

	fmt.Println(transferBit2Str("11101"))
	fmt.Println(transferNum2Bit(29, 16, "0123456789ABCDEF"))
	fmt.Println(transfer2To16("11101"))
}

func transfer2To16(str string) string {
	if len(str) < 1 {
		return ""
	}

	var (
		mapping string = "0123456789ABCDEF"
		size16  int    = 16
		rest    int    = len(str) % 4
		segmant int    = len(str) / 4
	)

	result := []string{}
	if rest > 0 {
		result = append(result, transferNum2Bit(transferBit2Str(str[0:rest]), size16, mapping))
	}

	if segmant > 0 {
		for i := 0; i < segmant; i++ {
			from := i*4 + rest
			result = append(result, transferNum2Bit(transferBit2Str(str[from:from+4]), size16, mapping))

		}
	}

	return strings.Join(result, "")

}

func transferNum2Bit(num int, size int, chars string) string {

	if len(chars) < size {
		panic("chars mapping length error")
	}

	result := []byte{}
	for num > 0 {
		res := num % size
		result = append(result, chars[res])
		num = num / size
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result)
}

func transferBit2Str(str string) int {
	offset := len(str) - 1
	result := 0
	for _, b := range str {
		if b == '1' {
			result += int(math.Pow(2, float64(offset)))
		}
		offset--
	}
	return result
}
