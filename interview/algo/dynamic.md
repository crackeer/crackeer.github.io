
## 最长递增子序列
- https://labuladong.gitee.io/algo/di-er-zhan-a01c6/dong-tai-g-a223e/dong-tai-g-6ea57/

```go
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
```