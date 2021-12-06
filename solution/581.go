package solution

/*
[2,6,4,8,10,9,15]
[1,2,3,4]
[1]
[2,1]
[1,3,5,4,2]
*/
func findUnsortedSubarray(nums []int) int {
	N := len(nums)
	if N <= 1 {
		return 0
	}
	pre := nums[0]
	p := -1
	for i := 1; i < N; i++ {
		if nums[i] < pre {
			p = i
			break
		}
		pre = nums[i]
	}
	if p == -1 {
		return 0
	}

	next := nums[N-1]
	q := -1
	for i := N - 2; i >= 0; i-- {
		if nums[i] > next {
			q = i
			break
		}
		next = nums[i]
	}
	if p == -1 {
		return 0
	}

	if p > q {
		p, q = q, p
	}

	min, max := MaxInt, MinInt
	for i := p; i <= q; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	p = p - 1
	for p >= 0 && nums[p] > min {
		p--
	}
	q = q + 1
	for q < N && nums[q] < max {
		q++
	}
	if p == -1 {
		p = p + 1
	}
	if q == N {
		q = q - 1
	}
	return q - p + 1
}
