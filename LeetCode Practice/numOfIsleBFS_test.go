package main

import "testing"

func TestNumIslandsBFS(t *testing.T) {
	tests := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			grid: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
	}

	for i, tt := range tests {
		got := numIslandsBFS(copyGridBFS(tt.grid))
		if got != tt.expected {
			t.Errorf("Test %d failed: expected %d, got %d", i+1, tt.expected, got)
		}
	}
}

// helper to deep copy a grid so tests don't interfere with each other
func copyGridBFS(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
