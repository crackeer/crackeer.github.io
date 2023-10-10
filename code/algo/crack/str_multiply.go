func multiply(str1 string, str2 string) string {
	if len(str1) < 1 || len(str2) < 1 {
		return "0"
	}

	str1Bytes := []byte(str1)
	str2Bytes := []byte(str2)
	mulResult := make([]int, len(str1Bytes)+len(str2Bytes))
	for i := len(str1Bytes) - 1; i >= 0; i-- {
		for j := len(str2Bytes) - 1; j >= 0; j-- {
			mul := int(str1Bytes[i]-'0') * int(str2Bytes[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := int(mul) + mulResult[p2]
			mulResult[p1] = sum / 10
			mulResult[p2] = sum % 10
		}
	}

	var (
		index int
	)
	for i, a := range mulResult {
		if a != 0 {
			index = i
			break
		}
	}
	result := []byte{}
	for i := index; i <= len(mulResult)-1; i++ {
		result = append(result, byte('0'+mulResult[i]))
	}
	fmt.Println(string(result))
	return string(result)
}