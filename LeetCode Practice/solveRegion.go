package main

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	mapped := make(map[[2]int]bool)

	var edgeRegion func(r, c int, currNei *[][2]int) bool
	edgeRegion = func(r, c int, currNei *[][2]int) bool {
		// out of bounds or not an 'O' → not part of region
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != 'O' {
			return false
		}
		// already visited
		if mapped[[2]int{r, c}] {
			return false
		}

		touchesEdge := (r == 0 || r == m-1 || c == 0 || c == n-1)

		mapped[[2]int{r, c}] = true
		*currNei = append(*currNei, [2]int{r, c})

		for _, d := range directions {
			if edgeRegion(r+d[0], c+d[1], currNei) {
				touchesEdge = true
			}
		}
		return touchesEdge
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !mapped[[2]int{i, j}] {
				var currentRegion [][2]int
				// If NOT touching border, flip it
				if !edgeRegion(i, j, &currentRegion) {
					for _, p := range currentRegion {
						r, c := p[0], p[1]
						board[r][c] = 'X'
					}
				}
			}
		}
	}
}
