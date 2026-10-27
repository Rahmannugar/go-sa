package main

// Quick sort is a divide and conquer algorithm that selects a pivot element from the array and partitions the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.
// Time complexity: O(n log n), worst-case time: O(n²)., Space complexity: O(log n)

// quickSort recursively sorts an array using the quick sort algorithm.
func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// partition partitions the array around a pivot element and returns the index of the pivot after partitioning.
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	arr := []int{8, 3, 1, 7, 0, 10, 2}

	quickSort(arr, 0, len(arr)-1)

	for _, value := range arr {
		println(value)
	}
}
