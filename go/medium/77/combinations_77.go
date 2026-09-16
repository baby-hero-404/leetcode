/**
 * https://leetcode.com/problems/combinations/description/
 */

package main

import "fmt"

func combine(n int, k int) [][]int {
	isTracking := make([]bool, n + 1)
    var result [][]int
	
	var backtrack func(from int, arr []int)
	backtrack = func(from int, arr []int) {
		if len(arr) == k {
			cloneArr := make([]int, k)
			copy(cloneArr, arr)
			result = append(result, cloneArr)
			return
		}
		for i:= from; i <= n; i++ {
			if !isTracking[i] {
				isTracking[i] = true
				newArr := append(arr, i)
				backtrack(i + 1,newArr)
				isTracking[i] = false
			}
		}
	}

	backtrack(1, nil)
	return result
}

func main() {
	results := combine(4,2)
	for _, result := range results {
		for _, num := range result {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}