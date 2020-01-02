package leetcode

import "sort"

func oddEven(xs []int) {
	i, j := 0, len(xs)-1
	for i != j {
		if xs[i]%2 != 0 {
			i++
		} else {
			xs[i], xs[j] = xs[j], xs[i]
			j--
		}
	}
	sort.Ints(xs[0 : i+1])
	sort.Ints(xs[i+1:])
}
