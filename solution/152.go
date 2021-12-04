package solution

func maxProduct(nums []int) int {
	type tuple struct {
		pos int
		neg int
	}
	if len(nums) == 1 {
		return nums[0]
	}

	pre := tuple{0, 0}
	res := 0
	for _, num := range nums {
		t := tuple{num, num}
		if num >= 0 {
			t.neg *= pre.neg
			if pre.pos > 0 {
				t.pos *= pre.pos
			}
		} else {
			t.pos *= pre.neg
			if pre.pos > 0 {
				t.neg *= pre.pos
			}
		}
		if t.pos > res {
			res = t.pos
		}
		pre = t
	}

	return res
}
