package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(("++++++"))
	result := plusString("993.456", "6.8")
	fmt.Println("Result=", result)
	fmt.Println("========")
}

// 问题：2个字符串数字相加
// 例如： “1.1" + "2.2" = "3.3""

// 不考虑负数、不考虑异常输入、输入肯定是小数（小数点肯定存在）

func plusString(num1, num2 string) string {
	a1, a2 := cal(num1)
	b1, b2 := cal(num2)

	bytes1 := []byte(num1)
	bytes2 := []byte(num2)
	if a2 > b2 {
		bytes2 = appendChar(bytes2, '0', a2-b2)
	} else if b2 > a2 {
		bytes1 = appendChar(bytes1, '0', b2-a2)
	}

	if a1 > b1 {
		bytes2 = preAppendChar(bytes2, '0', a1-b1)
	} else if b1 > a1 {
		bytes1 = preAppendChar(bytes1, '0', b1-a1)
	}

	result := make([]byte, len(bytes1)+1)

	var (
		extra uint
		c     byte
		index int = len(result) - 1
	)

	fmt.Println(string(bytes1), string(bytes2))

	for i := len(bytes1) - 1; i >= 0; i-- {
		if bytes1[i] == '.' {
			result[index] = '.'
			index--
			continue
		}
		extra, c = plusChar(bytes1[i], bytes2[i], extra)
		fmt.Println(extra)
		result[index] = c
		index--
		fmt.Println(index)
	}
	fmt.Println(string(result), "See")
	if extra > 0 {
		result[index] = '1'
	}

	return strings.Trim(string(result), "0")

}

func plusChar(a, b byte, extra uint) (uint, byte) {

	sum := uint(a-'0'+(b-'0')) + extra
	if sum >= 10 {
		return 1, byte(sum%10 + '0')
	}
	return 0, byte(sum%10 + '0')
}

func preAppendChar(bytes []byte, c byte, size int) []byte {
	result := []byte{}
	for i := 0; i < size; i++ {
		result = append(result, c)
	}
	result = append(result, bytes...)
	return result
}

func appendChar(bytes []byte, c byte, size int) []byte {
	for i := 0; i < size; i++ {
		bytes = append(bytes, c)
	}
	return bytes
}

func cal(num1 string) (int, int) {
	parts := strings.Split(num1, ".")
	return len(parts[0]), len(parts[1])
}
