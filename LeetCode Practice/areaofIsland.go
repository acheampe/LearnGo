package main

func maxAreaOfIsland(grid [][]int) int {

	var m, n, maxArea int

	m = len(grid)

	if m > 0 {
		n = len(grid[0])
	}

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var currAreaDFS func(r, c int) int

	currAreaDFS = func(r, c int) int {

		// Base Case
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0
		area := 1

		for _, d := range directions {
			dr, dc := d[0]+r, d[1]+c
			area += currAreaDFS(dr, dc)
		}

		return area

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == 1 {
				maxArea = max(maxArea, currAreaDFS(i, j))
			}
		}
	}

	return maxArea
}
