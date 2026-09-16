/**
 * https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/description/
 */

package main

import "fmt"

func main() {
	print("start test 1\n")
	var grid = [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}
	k := 1
	result := shortestPath(grid, k)
	fmt.Println(result)

	print("start test 2\n")
	grid = [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}
	k = 1
	result = shortestPath(grid, k)
	fmt.Println(result)
}

type State struct {
	r, c, k int
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	if m == 1 && n == 1 {
		return 0
	}
	if k >= m+n-2 {
		return m + n - 2
	}
	maxRemainingK := make([]int, m*n)
	for i := range maxRemainingK {
		maxRemainingK[i] = -1 // -1 means not visited
	}
	maxRemainingK[0] = k

	moves := [][2]int{
		{1, 0},  // down
		{-1, 0}, // up
		{0, 1},  // right
		{0, -1}, // left
	}
	queue := []State{{0, 0, k}}
	head, steps := 0, 0

	for head < len(queue) {
		end := len(queue)

		for head < end {
			cur := queue[head]
			head++

			if cur.r == m-1 && cur.c == n-1 {
				return steps
			}

			for _, d := range moves {
				r, c := cur.r+d[0], cur.c+d[1]

				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}

				remaining := cur.k - grid[r][c]

				if remaining < 0 || maxRemainingK[r*n+c] >= remaining {
					continue
				}

				maxRemainingK[r*n+c] = remaining
				queue = append(queue, State{r, c, remaining})
			}
		}

		steps++
	}

	return -1
}
