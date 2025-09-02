package main

import "container/list"

func maxAreaOfIsland(grid [][]int) int {

	m := len(grid)
	var n int

	if m > 0 {
		n = len(grid[0])
	} else {
		return 0
	}

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	queue := list.New()
	var maxArea int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == 1 {
				currentArea := 0

				queue.PushBack([]int{i, j}) // enqueue
				grid[i][j] = 0

				for queue.Len() > 0 {
					e := queue.Front()
					curr := e.Value.([]int)
					queue.Remove(e) // pop curr node
					r, c := curr[0], curr[1]
					currentArea++

					for _, d := range directions {
						dr, dc := r+d[0], c+d[1]

						if 0 <= dr && dr < m && 0 <= dc && dc < n && grid[dr][dc] == 1 {
							queue.PushBack([]int{dr, dc})
							grid[dr][dc] = 0

						}
					}
				}

				if currentArea > maxArea {
					maxArea = currentArea
				}

			}
		}
	}

	return maxArea
}
