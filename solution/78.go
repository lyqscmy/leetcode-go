package solution

func subsets(nums []int) [][]int {
	N := len(nums)
	M := (1 << N) - 1
	res := make([][]int, 0, 1<<N)
	for i := 0; i <= M; i++ {
		xs := make([]int, 0)
		for j := 0; j < N; j++ {
			mask := 1 << j
			if i&mask > 0 {
				xs = append(xs, nums[j])
			}
		}
		res = append(res, xs)
	}
	return res
}
