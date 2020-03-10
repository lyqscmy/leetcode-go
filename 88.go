package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	copy(nums1[n:], nums1[:m])
	// 合并排序，i是nums1的游标，j是nums2的游标，k是写入下一个值的游标
	for i, j, k := n, 0, 0; j < n; k++ {
		if i >= m+n {
			copy(nums1[k:], nums2[j:])
			break
		}

		if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}
