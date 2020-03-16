package solution

func lengthOfLIS(nums []int) int {
	N := len(nums)
	tails := make([]int, N)
	res := 0

	for _, num := range nums {
		i, j := 0, res
		for i < j {
			h := int(uint(i+j) >> 1)
			if tails[h] < num {
				i = h + 1
			} else {
				j = h
			}
		}
		if i < res {
			tails[i] = num
		} else {
			tails[res] = num
			res++
		}
	}

	return res
}
