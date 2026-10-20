package main

import "math"

// o(n): it loops through the array once to find the minimum value, so the time complexity is linear O(n). The space complexity is constant O(1) because it uses a fixed amount of space regardless of the input size.
// findMax finds the maximum value in an array of integers.
// Time O(n), space O(1).
func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0 // Return 0 for empty array
	}
	max := math.Inf(-1)

	for _, num := range arr {
		if float64(num) > max {
			max = float64(num)
		}
	}
	return int(max)
}

// o(n^2): it uses nested loops to compare each first name with each last name, resulting in a quadratic time complexity O(n^2). The space complexity is linear O(n) because it creates a new slice to store the combined names, which grows with the input size.
// findName finds a full name in slices of first and last names.
// Time O(n^2), space O(n).
func findName(firstNames, lastNames []string, fullName string) bool {
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			if firstName+" "+lastName == fullName {
				return true
			}
		}
	}
	return false
}

// o(1): it performs a constant-time lookup in the map, so the time complexity is O(1). The space complexity is also O(1) because it uses a fixed amount of space regardless of the input size.
// findLastName finds the last name associated with a given first name in a map of names.
// Time O(1), space O(1).
func findLastName(namesDictionary map[string]string, firstName string) string {
	var lastName string
	lastName = namesDictionary[firstName]
	return lastName
}

// o(log n): it uses a binary search algorithm to find the index of the target value in the sorted array, resulting in a logarithmic time complexity O(log n). The space complexity is constant O(1) because it uses a fixed amount of space regardless of the input size.
// binarySearch finds the index of a target value in a sorted array of integers using binary search.
// Time O(log n), space O(1).
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Target not found
}

func main() {
	//find maximum value
	arr := []int{5, 2, 8, 1, 9}
	max := findMax(arr)
	println(max)

	//find full name
	firstNames := []string{"Abdulrahmon", "Jane", "Alice"}
	lastNames := []string{"Adenuga", "Smith", "Johnson"}
	fullName := "Abdulrahmon Adenuga"
	found := findName(firstNames, lastNames, fullName)
	println(found)

	//find last name
	namesDictionary := map[string]string{
		"Abdulrahmon": "Adenuga",
		"Jane":        "Smith",
		"Alice":       "Johnson",
	}
	firstName := "Abdulrahmon"
	lastName := findLastName(namesDictionary, firstName)
	println(lastName)

	//binary search
	sortedArr := []int{1, 2, 3, 4, 5}
	target := 3
	index := binarySearch(sortedArr, target)
	println(index)
}
