package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	numerator := b*b - 2*a*c
	denominator := a * a

	divisor := gcd(numerator, denominator)
	u := numerator / divisor
	v := denominator / divisor

	fmt.Println(u, v)
}
