package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	totalSum := 0
	for _, val := range a {
		totalSum += val
	}

	count := 0
	prefixSum := 0

	for i := 0; i < n-1; i++ {
		prefixSum += a[i]
		suffixSum := totalSum - prefixSum
		if prefixSum == suffixSum {
			count++
		}
	}

	fmt.Println(count)
}
