package main

// Insertion sort is a simple sorting algorithm that builds the final sorted array one item at a time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort but has the advantage of being simple to implement and efficient for small data sets or nearly sorted data.
// Time complexity: O(n^2), Space complexity: O(1)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}

	return arr
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	sortedArr := insertionSort(arr)
	for _, value := range sortedArr {
		println(value)
	}
}
