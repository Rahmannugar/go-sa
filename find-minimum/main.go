package main

import "math"

// findMinimum finds the minimum value in an array of integers.
// Time O(n), space O(1).

func findMinimum(arr []int) int {

	if len(arr) == 0 {
		return 0 // Return 0 for empty array
	}

	var min = math.Inf(1)

	for _, num := range arr {
		if float64(num) < min {
			min = float64(num)
		}
	}
	return int(min)
}

func main() {
	arr := []int{5, 2, 8, 1, 9}
	min := findMinimum(arr)
	println(min)
}
