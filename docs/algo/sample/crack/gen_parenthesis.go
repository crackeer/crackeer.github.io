// 括号对生成
func generateParenthesis(count int) []string {
	list := &[]string{}
	backtrace1(count, count, []byte{}, list)
	for _, v := range *list {
		fmt.Println(v)
	}
	return *list
}

func backtrace1(left, right int, current []byte, list *[]string) {
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		*list = append(*list, string(current))
		return
	}

	current = append(current, '(')
	backtrace1(left-1, right, current, list)
	current = current[0 : len(current)-1]

	current = append(current, ')')
	backtrace1(left, right-1, current, list)
	current = current[0 : len(current)-1]
}