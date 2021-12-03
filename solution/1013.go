package solution

// [0:i), [i:j), [j:N) 任一都不为空
func canThreePartsEqualSum(A []int) bool {
	N := len(A)
	if N < 3 {
		return false
	}
	total := sum(A)
	if (total % 3) != 0 {
		return false
	}

	part := total / 3
	i := 0
	acc := 0
	for ; i < N; i++ {
		acc += A[i]
		if acc == part {
			break
		}
	}

	acc = 0
	j := i + 1
	for ; j < N; j++ {
		acc += A[j]
		if acc == part {
			break
		}
	}
	// j<N
	return j < N-1
}
