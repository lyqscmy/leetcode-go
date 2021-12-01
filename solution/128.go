package solution

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	res := 0
	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}
		count := 1
		for {
			if _, ok := set[num+1]; !ok {
				break
			}
			num++
			count++
		}
		if count > res {
			res = count
		}
	}
	return res
}
