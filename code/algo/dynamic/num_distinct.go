func numDistinct(s string, t string) int {

	if len(s) < len(t) {
		return 0
	}

	if len(t) < 1 {
		return 1
	}

	if s == t {
		return 1
	}
	result := 0
	for i, b := range []byte(s) {
		if b == t[0] {
			result += numDistinct(s[i+1:], t[1:])
		}
	}
	return result
}
