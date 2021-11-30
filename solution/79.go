package solution

func exist(board [][]byte, word string) bool {
	path = make(map[Point]struct{})
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
		path[a] = struct{}{}
		if dfs(m, word[1:], a) {
			return true
		}
		delete(path, a)
	}
	return false
}

type Point struct {
	x int
	y int
}

var path map[Point]struct{}

// len(m) <= len(word)
func dfs(m map[byte][]Point, word string, pre Point) bool {
	if len(word) == 0 {
		return true
	}
	A := m[word[0]]
	for _, a := range A {
		if _, ok := path[a]; ok {
			continue
		}
		if adjacent(pre, a) {
			path[a] = struct{}{}

			if dfs(m, word[1:], a) {
				return true
			}
			delete(path, a)
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
