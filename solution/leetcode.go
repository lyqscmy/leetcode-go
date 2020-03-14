package solution

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
	N := len(h)
	d := make([]int, N)
	for i := 0; i < N; i++ {
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
