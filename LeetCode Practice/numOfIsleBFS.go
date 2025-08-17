package main

import (
	"container/list"
)

func numIslandsBFS(grid [][]byte) int {

	var count, m, n int

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

	queue := list.New()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0'
				// map island
				queue.PushBack([]int{i, j})

				for queue.Len() > 0 {
					e := queue.Front()
					curr := e.Value.([]int)
					queue.Remove(e)

					r, c := curr[0], curr[1]
					for _, d := range directions {
						nr, nc := r+d[0], c+d[1]

						if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == '1' {
							queue.PushBack([]int{nr, nc})
							grid[nr][nc] = '0'
						}
					}

				}
			}
		}
	}

	return count
}
