package main

import (
	"fmt"
)

func atoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= 999; i++ {
		s := fmt.Sprint(i)
		doubleStr := s + s
		num := atoi(doubleStr)
		if num <= n {
			count++
		}
	}
	fmt.Println(count)
}
