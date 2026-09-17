/**
 * https://leetcode.com/problems/can-place-flowers/description/
 */

package main

import "fmt"

func main() {
	print("start test 1\n")
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	result := canPlaceFlowers(flowerbed, n)
	fmt.Println(result)

	print("start test 2\n")
	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2
	result = canPlaceFlowers(flowerbed, n)
	fmt.Println(result)

	print("start test 3\n")
	flowerbed = []int{1, 0, 0, 0, 0, 1}
	n = 2
	result = canPlaceFlowers(flowerbed, n)
	fmt.Println(result)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := range flowerbed {
		if flowerbed[i] != 0 {
			continue
		}

		if (i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {

			flowerbed[i] = 1
			n--

			if n == 0 {
				return true
			}
		}
	}

	return false
}
