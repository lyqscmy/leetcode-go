package solution

func maximalSquare(matrix [][]byte) int {
	dp1 := make([]int, len(matrix[0])+1)
	dp2 := make([]int, len(matrix[0])+1)
	res := 0
	for _, row := range matrix {
		for j, v := range row {
			if v == 0x31 {
				dp2[j+1] = min(min(dp1[j], dp1[j+1]), dp2[j]) + 1
				res = max(dp2[j+1], res)
			}
		}
		dp1, dp2 = dp2, dp1
		for i := 0; i < len(dp2); i++ {
			dp2[i] = 0
		}
	}
	return res * res
}
