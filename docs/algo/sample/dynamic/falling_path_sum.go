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
