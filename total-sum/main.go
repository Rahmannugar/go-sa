package main

// findTotalSum calculates the total sum of an array of integers.
// Time O(n), space O(1).

func findTotalSum(arr []int) int {
	sum := 0

	if len(arr) == 0 {
		return 0 // Return 0 for empty array
	}

	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	arr := []int{5, 2, 8, 1, 9}
	arr2 := []int{}
	total := findTotalSum(arr)
	totalEmpty := findTotalSum(arr2)
	println(total)
	println(totalEmpty)
}
