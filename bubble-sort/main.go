package main

//why would you even use this, it is so impractical lmao, anyways.

// Bubble sort is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted.
// Time complexity: O(n^2), Space complexity: O(1)
func bubbleSort(arr []int) []int {
	swapping := true
	end := len(arr)

	for swapping {
		swapping = false
		for i := 1; i < end; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapping = true
			}
		}
		end--
	}
	return arr
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	sortedArr := bubbleSort(arr)
	for _, value := range sortedArr {
		println(value)
	}
}
