package solution

type Vertex struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	M = len(grid)
	N = len(grid[0])
	rot := make([][]int, M)
	for i := 0; i < M; i++ {
		rot[i] = make([]int, N)
	}

	noFresh := true
	q := NewQueue()
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				rot[i][j] = -1
			} else if grid[i][j] == 1 {
				noFresh = false
				rot[i][j] = MaxInt
			} else {
				q.Add(Vertex{i, j})
			}
		}
	}
	if noFresh {
		return 0
	}

	// bfs
	dir := []Vertex{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	for q.Length() > 0 {
		v := q.Remove().(Vertex)
		for _, d := range dir {
			gray := rotting(rot, v, d)
			if gray {
				q.Add(Vertex{v.x + d.x, v.y + d.y})
			}
		}
	}

	// find the max rotted time
	res := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			t := rot[i][j]
			if t > res {
				res = t
			}
		}
	}
	if res == MaxInt {
		return -1
	}
	return res
}

func rotting(rot [][]int, v Vertex, d Vertex) bool {
	a := rot[v.x][v.y]
	i := v.x + d.x
	if i < 0 || i >= M {
		return false
	}
	j := v.y + d.y
	if j < 0 || j >= N {
		return false
	}
	b := rot[i][j]
	if b != -1 && a+1 < b {
		rot[i][j] = a + 1
	}
	return b == MaxInt
}
