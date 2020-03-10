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
	for i, j, k := n, 0, 0; j < n; k++ {
		// fmt.Println(nums1[:k], nums1[i:], nums2[j:])
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

func reverse(nums1 []int, from int, to int) {
	for from < to {
		nums1[from], nums1[to] = nums1[to], nums1[from]
		from++
		to--
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for i, j := m-1, 0; i >= 0 && j < n; {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
		} else {
			i--
		}
	}
	return false
}

func majorityElement(nums []int) int {
	res := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			res, count = num, 0
		}
		if num != res {
			count--
		} else {
			count++
		}
	}
	return res
}

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

// // left to right = top to bottom
// func dtoh(d []int) []int {
// 	n := len(d)
// 	h := NewDeque(n)
// 	for i := 0; i < n; i++ {
// 		// step 2
// 		if h.Len() != 0 {
// 			h.AddFirst(h.RemoveLast())
// 		}

// 		// step 1
// 		h.AddFirst(d[i])
// 	}
// 	return h.Items()
// }

// left to right = top to bottom
func htod(h []int) []int {
	n := len(h)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		// step 1
		d[i] = h[0]
		h = h[1:]

		// step 2
		if len(h) != 0 {
			x := h[0]
			h = h[1:]
			h = append(h, x)
		}
	}
	return d
}

// Compare do
func Compare(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ParseInt do
func ParseInt(s string) int32 {
	var n int32 = 0
	var cutoff int32 = (1<<32)/10 + 1
	for i := 0; i < len(s); i++ {
		if n >= cutoff {
			return -1
		}
		n *= 10
		n1 := n + int32(s[i]-'0')
		if n1 < n {
			return -1
		}
		n = n1
	}
	return n
}

func addList(l1 *ListNode, l2 *ListNode) *ListNode {
	dumy := &ListNode{}
	node := dumy
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		node = node.Next
		carry /= 10
	}
	if carry != 0 {
		node.Next = &ListNode{carry, nil}
	}
	return dumy.Next
}
