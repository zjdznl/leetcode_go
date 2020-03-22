package leetcode

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	expandGrid := make([][]int, m*3)
	visited := make([][]int8, m*3)
	for i := range expandGrid {
		expandGrid[i] = make([]int, n*3)
		visited[i] = make([]int8, n*3)
	}

	for i, v := range grid {
		for j := range v {
			i3 := i * 3
			j3 := j * 3
			switch grid[i][j] {
			case 1:
				expandGrid[i3+1][j3] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3+1][j3+2] = 1
			case 2:
				expandGrid[i3][j3+1] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3+2][j3+1] = 1
			case 3:
				expandGrid[i3+1][j3] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3+2][j3+1] = 1
			case 4:
				expandGrid[i3+2][j3+1] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3+1][j3+2] = 1
			case 5:
				expandGrid[i3+1][j3] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3][j3+1] = 1
			case 6:
				expandGrid[i3][j3+1] = 1
				expandGrid[i3+1][j3+1] = 1
				expandGrid[i3+1][j3+2] = 1
			default:
				panic("not possible")

			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if expandGrid[i][j] == 1 {
				return dfs1391(expandGrid, visited, i, j)
			}
		}
	}
	return false
}

func dfs1391(expandGrid [][]int, visited [][]int8, i, j int) bool {
	if i < 0 || j < 0 || i >= len(expandGrid) || j >= len(expandGrid[0]) || visited[i][j] == 1 || expandGrid[i][j] == 0 {
		return false
	}

	visited[i][j] = 1
	if i > len(expandGrid)-3 && j > len(expandGrid[0])-3 {
		return true
	}
	//println("current: i: ", i, ", j: ", j)
	//println("next: i: ", i, ", j: ", j-1)
	//println("next: i: ", i, ", j: ", j+1)
	//println("next: i: ", i-1, ", j: ", j)
	//println("next: i: ", i+1, ", j: ", j)

	// 上下左右
	return dfs1391(expandGrid, visited, i, j-1) ||
		dfs1391(expandGrid, visited, i, j+1) ||
		dfs1391(expandGrid, visited, i-1, j) ||
		dfs1391(expandGrid, visited, i+1, j)
}
