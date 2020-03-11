package solution

/**
 * 1. a ^ a = 0 归零律
 * 2. a ^ 0 = a 恒等律
 * 3. a ^ b ^ c = a ^ c ^ b 交换律
 */
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
