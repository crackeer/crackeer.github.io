/// [subSequence]
func subSequenceLen(list []int) int {
	dp := make([]int, len(list))
	dp[0] = 1

	for i := 1; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[j] > list[i] {
				dp[i] = util.Max(dp[i-1], dp[j]+1)
			}
		}
	}
	return util.Max(dp...)
}

/// [subSequence]

/// [coin-change]
func coinChange(coins []int, amount int) int {

	if amount < 0 {
		return -1
	}

	if amount < 1 {
		return 0
	}
	table := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		table[i] = math.MaxInt
		for _, c := range coins {
			if i-c >= 0 {
				if table[i-c] != math.MaxInt {
					table[i] = min(table[i], table[i-c]+1)
				}
			}
		}
	}
	if table[amount] == math.MaxInt {
		return -1
	}
	return table[amount]
}

/// [coin-change]

/// [minFallingPathSum]
func minFallingPathSum(matrix [][]int) int {
	data := minWays(matrix)
	result := math.MaxInt
	for _, item := range data {
		result = min(result, item)
	}
	return result
}

func minWays(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	}

	if len(matrix) == 1 {
		return matrix[0]
	}

	result := make([]int, len(matrix[0]))
	minSums := minWays(matrix[1:])
	for index, val := range matrix[0] {
		result[index] = val + minSums[index]
		if index-1 >= 0 {
			result[index] = min(result[index], val+minSums[index-1])
		}

		if index+1 <= len(matrix[0]) {
			result[index] = min(result[index], val+minSums[index+1])
		}
	}
	return result
}

/// [minFallingPathSum]

/// [numDistinct]

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

/// [numDistinct]

/// [getFullOrder]

func getFullOrder(str string) []string {
	if len(str) < 1 {
		return nil
	}
	if len(str) < 2 {
		return []string{str}
	}
	bytes := []byte(str)
	result := []string{}
	for i := 0; i < len(bytes); i++ {
		if i-1 > -1 && bytes[i] == bytes[i-1] {
			continue
		}
		bytes[0], bytes[i] = bytes[i], bytes[0]
		orders := getFullOrder(string(bytes[1:]))
		for _, item := range orders {
			tmp := make([]byte, len(item)+1)

			copy(tmp, bytes[0:1])
			copy(tmp[1:], item)
			result = append(result, string(tmp))
		}
		bytes[i], bytes[0] = bytes[0], bytes[i]
	}
	return result
}

/// [getFullOrder]