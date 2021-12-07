package solution

func exist(board [][]byte, word string) bool {
	M = len(board)
	N = len(board[0])
	s := make(map[byte]struct{})
	for i := 0; i < len(word); i++ {
		s[word[i]] = struct{}{}
	}

	m := make(map[byte][]Point)
	for i, row := range board {
		for j, ch := range row {
			if _, ok := s[ch]; ok {
				m[ch] = append(m[ch], Point{i, j})
			}
		}
	}

	A := m[word[0]]
	for _, a := range A {
		SetVisited(a.x, a.y)
		if dfs(m, word[1:], a) {
			return true
		}
		UnSetVisited(a.x, a.y)
	}
	return false
}

// len(m) <= len(word)
func dfs(m map[byte][]Point, word string, pre Point) bool {
	if len(word) == 0 {
		return true
	}
	A := m[word[0]]
	for _, a := range A {
		if IsVisited(a.x, a.y) {
			continue
		}
		if adjacent(pre, a) {

			SetVisited(a.x, a.y)
			if dfs(m, word[1:], a) {
				return true
			}
			UnSetVisited(a.x, a.y)
		}
	}
	return false
}

func adjacent(a Point, b Point) bool {
	return Square(a.x-b.x)+Square(a.y-b.y) == 1
}

func Square(x int) int {
	return x * x
}
