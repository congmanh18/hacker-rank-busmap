package main

import (
	"fmt"
)

func countFactor(x, f int) int {
	if x == 0 {
		return 0
	}
	count := 0
	for x%f == 0 {
		count++
		x /= f
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	count2 := make([]int, n)
	count5 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		count2[i] = countFactor(a[i], 2)
		count5[i] = countFactor(a[i], 5)
	}

	maxZeros := 0
	sum2, sum5 := 0, 0

	for right := 0; right < n; right++ {
		if a[right] == 0 {
			sum2, sum5 = 0, 0
			continue
		}
		sum2 += count2[right]
		sum5 += count5[right]

		maxZeros = max(maxZeros, min(sum2, sum5))
	}

	fmt.Println(maxZeros)
}
