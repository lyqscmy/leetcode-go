package solution

/*
[[1,2,1,2,1,2],[2,2,2,2,1,2],[1,2,2,2,1,2]]
1
*/
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	M = len(grid)
	N = len(grid[0])
	visited = make([][]bool, M)
	OriginColor := grid[row][col]
	border := make([]Point, 0)
	q := NewQueue()
	q.Add(Point{row, col})

	for q.Length() > 0 {
		cell := q.Remove().(Point)
		SetVisited(cell.x, cell.y)
		surround := 0
		for _, d := range dirs {
			i := cell.x + d.x
			if i < 0 || i >= M {
				continue
			}
			j := cell.y + d.y
			if j < 0 || j >= N {
				continue
			}
			if grid[i][j] != OriginColor {
				continue
			}
			surround++
			if IsVisited(i, j) {
				continue
			}
			q.Add(Point{i, j})
		}
		if surround != 4 {
			border = append(border, cell)
		}
	}
	for _, p := range border {
		grid[p.x][p.y] = color
	}
	return grid
}
