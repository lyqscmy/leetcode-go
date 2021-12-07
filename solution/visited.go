package solution

var visited [][]bool

func SetVisited(x, y int) {
	if visited[x] == nil {
		visited[x] = make([]bool, N)
	}
	visited[x][y] = true
}

func UnSetVisited(x, y int) {
	visited[x][y] = false
}

func IsVisited(x, y int) bool {
	if visited[x] == nil {
		return false
	}
	return visited[x][y]
}
