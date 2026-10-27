package main

// Merge sort is a divide and conquer algorithm that divides the input array into two halves, calls itself for the two halves, and then merges the two sorted halves.
// Time complexity: O(n log n), Space complexity: O(n)

// mergeSort recursively sorts an array using the merge sort algorithm.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice.
func merge(first, second []int) []int {
	result := make([]int, 0, len(first)+len(second))
	i, j := 0, 0

	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			result = append(result, first[i])
			i++
		} else {
			result = append(result, second[j])
			j++
		}
	}

	for i < len(first) {
		result = append(result, first[i])
		i++
	}

	for j < len(second) {
		result = append(result, second[j])
		j++
	}
	return result
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10, 5, 24, 7, 30, 14, 50, 20, 72}
	sortedArr := mergeSort(arr)
	for _, value := range sortedArr {
		println(value)
	}
}
