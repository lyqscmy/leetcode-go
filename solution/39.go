package solution

func combinationSum(candidates []int, target int) [][]int {
	s := &state{}
	for i := 0; i < len(candidates); i++ {
		s.dfs(candidates[i:], target)
	}
	return s.res
}

type state struct {
	res  [][]int
	path []int
}

func (s *state) dfs(candidates []int, target int) {
	x := candidates[0]
	s.path = append(s.path, x)
	target -= x
	defer func() {
		s.path = s.path[:len(s.path)-1]
	}()
	if target == 0 {
		s.res = append(s.res, clone(s.path))
		return
	}
	if target < 0 {
		return
	}
	N := len(candidates)
	for i := 0; i < N; i++ {
		s.dfs(candidates[i:], target)
	}
	return
}

func clone(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
