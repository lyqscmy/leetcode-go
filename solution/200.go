package solution

type IslandWalker struct {
	M    int
	N    int
	grid [][]byte
	res  int
}

func numIslands(grid [][]byte) int {
	w := NewIslandWalker(grid)
	for m, row := range grid {
		for n, v := range row {
			if v == 0x31 {
				w.res++
				w.dfs(m, n)
			}
		}
	}
	return w.res
}

func NewIslandWalker(grid [][]byte) *IslandWalker {
	return &IslandWalker{grid: grid, M: len(grid), N: len(grid[0])}
}

func (w *IslandWalker) dfs(m, n int) {
	w.grid[m][n] = 0x30
	// down
	if m+1 < w.M && w.grid[m+1][n] == 0x31 {
		w.dfs(m+1, n)
	}

	// up
	if m-1 >= 0 && w.grid[m-1][n] == 0x31 {
		w.dfs(m-1, n)
	}

	// right
	if n+1 < w.N && w.grid[m][n+1] == 0x31 {
		w.dfs(m, n+1)
	}

	// left
	if n-1 >= 0 && w.grid[m][n-1] == 0x31 {
		w.dfs(m, n-1)
	}
}
