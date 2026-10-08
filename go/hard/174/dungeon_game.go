/**
 * https://leetcode.com/problems/dungeon-game/description/
 */

package main

import "fmt"

func main() {
	fmt.Println("start test 1")

	dungeon := [][]int{
		{0},
	}

	result := calculateMinimumHP(dungeon)
	assertEqual(1, result)

	fmt.Println("start test 2")

	dungeon = [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}

	result = calculateMinimumHP(dungeon)
	assertEqual(7, result)

	fmt.Println("start test 3")

	dungeon = [][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}

	result = calculateMinimumHP(dungeon)
	assertEqual(3, result)
}

func assertEqual(expected, actual int) {
	if expected != actual {
		fmt.Printf("FAIL: expected=%d, got=%d\n", expected, actual)
		return
	}

	fmt.Printf("PASS: expected=%d, got=%d\n", expected, actual)
}

var maxHPValue = 200*200*1000 + 1

func calculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])

	// requiredHP[col] = minimum HP required to enter the current cell and survive until the destination.
	requiredHP := make([]int, cols)

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {

			// The destination cell.
			if row == rows-1 && col == cols-1 {
				requiredHP[col] = max(1, 1-dungeon[row][col])
				continue
			}

			// Find the minimum HP required by the next cell.
			minNextHP := maxHPValue

			// Move down.
			if row+1 < rows {
				minNextHP = min(minNextHP, requiredHP[col])
			}

			// Move right.
			if col+1 < cols {
				minNextHP = min(minNextHP, requiredHP[col+1])
			}

			// Calculate the HP required before entering the current cell.
			neededHP := minNextHP - dungeon[row][col]

			// HP must always be at least 1.
			requiredHP[col] = max(1, neededHP)
		}
	}

	return requiredHP[0]
}
