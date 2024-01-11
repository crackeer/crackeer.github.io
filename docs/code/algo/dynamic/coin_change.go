func coinChangeX(coins []int, amount int) int {

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