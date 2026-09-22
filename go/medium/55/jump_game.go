/**
 * https://leetcode.com/problems/jump-game/description/
 */

package main

func main() {
	println("start test 1")
	results := canJump([]int{2, 3, 1, 1, 4})
	println(results)

	println("start test 2")
	results = canJump([]int{3, 2, 1, 0, 4})
	println(results)
}

func canJump(nums []int) bool {
	maxReach := 0

	for i, jump := range nums {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+jump)

		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return true
}
