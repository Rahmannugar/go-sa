package main

import "math"

//getExponent calculates the exponent of an array of integers based on the average and length of the array.
// Time O(n), space O(1).

func getExponent(arr []int) {
	arrLength := len(arr)

	if arrLength == 0 {
		println("Array is empty")
		return
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	avg := float64(sum) / float64(arrLength)
	exponent := math.Round(avg * math.Pow(float64(arrLength), 1.2))
	println("Exponent:", exponent)
}

func main() {
	arr := []int{12, 12, 12}
	arr2 := []int{7, 4, 3, 100, 765, 2344, 1, 2, 32}
	getExponent(arr)
	getExponent(arr2)
}
