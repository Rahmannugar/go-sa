package main

import "math"

//getLogScore calculates the logarithmic score based on the count and average percentage.
// Time O(1), space O(1).

func getLogScore(count int, averagePercent float64) float64 {
	return math.Round(averagePercent * math.Log2(float64(count)))
}

func main() {
	count := 40000
	averagePercent := 0.3
	result := getLogScore(count, averagePercent)
	println("Score:", result)
}
