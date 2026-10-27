package main

import "fmt"

// fib calculates the nth Fibonacci number by adding the previous two numbers.
// Time complexity: O(n), Space complexity: O(1)
func fib(n int) int {
	if n < 0 {
		return -1 // Invalid input
	}

	if n <= 1 {
		return n
	}

	previous := 0
	current := 1

	for i := 2; i <= n; i++ {
		next := previous + current
		previous = current
		current = next
	}

	return current
}

// powerSet recursively generates every subset by including or excluding each element.
// Time complexity: O(n * 2^n), Space complexity: O(n * 2^n)
func powerSet(inputSet []int) [][]int {
	if len(inputSet) == 0 {
		return [][]int{{}}
	}

	first := inputSet[0]
	remaining := inputSet[1:]

	remainingSubsets := powerSet(remaining)
	subsets := [][]int{}

	for _, subset := range remainingSubsets {
		subsets = append(subsets, subset)

		withFirst := append([]int{first}, subset...)
		subsets = append(subsets, withFirst)
	}

	return subsets
}

func main() {
	println(fib(10))
	fmt.Println(powerSet([]int{1, 2, 3}))
}
