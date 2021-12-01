package solution

var N int

func exist(board [][]byte, word string) bool {
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

	visited = make([][]bool, len(board))
	A := m[word[0]]
	for _, a := range A {
		setVisited(a, visited)
		if dfs(m, word[1:], a) {
			return true
		}
		unSetVisited(a, visited)
	}
	return false
}

func setVisited(p Point, visited [][]bool) {
	if visited[p.x] == nil {
		visited[p.x] = make([]bool, N)
	}
	visited[p.x][p.y] = true
}

func unSetVisited(p Point, visited [][]bool) {
	visited[p.x][p.y] = false
}
func isVisited(p Point, visited [][]bool) bool {
	if visited[p.x] == nil {
		return false
	}
	return visited[p.x][p.y]
}

type Point struct {
	x int
	y int
}

var visited [][]bool

// len(m) <= len(word)
func dfs(m map[byte][]Point, word string, pre Point) bool {
	if len(word) == 0 {
		return true
	}
	A := m[word[0]]
	for _, a := range A {
		if isVisited(a, visited) {
			continue
		}
		if adjacent(pre, a) {

			setVisited(a, visited)
			if dfs(m, word[1:], a) {
				return true
			}
			unSetVisited(a, visited)
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
