/**
 * https://leetcode.com/problems/24-game/description/
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	print("start test 1\n")
	cards := []int{4, 1, 8, 7}
	result := judgePoint24(cards)
	fmt.Println(result)

	print("start test 2\n")
	cards = []int{1, 2, 1, 2}
	result = judgePoint24(cards)
	fmt.Println(result)
}

func judgePoint24(cards []int) bool {
	nums := make([]float64, len(cards))

	for i, card := range cards {
		nums[i] = float64(card)
	}

	return dfs(nums)
}

func dfs(nums []float64) bool {
	// Only one number remains.
	if len(nums) == 1 {
		return math.Abs(nums[0]-24) < 1e-6
	}

	// Try every pair of numbers.
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			// Numbers that are not selected.
			remaining := []float64{}

			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					remaining = append(remaining, nums[k])
				}
			}

			// Try all possible results from nums[i] and nums[j].
			a, b := nums[i], nums[j]
			results := []float64{
				a + b,
				a - b,
				b - a,
				a * b,
			}

			if math.Abs(b) > 1e-6 {
				results = append(results, a/b)
			}

			if math.Abs(a) > 1e-6 {
				results = append(results, b/a)
			}

			for _, result := range results {
				if dfs(append(remaining, result)) {
					return true
				}
			}
		}
	}

	return false
}
