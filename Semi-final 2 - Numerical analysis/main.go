package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[0][0] = 1

	for pos := 0; pos < K; pos += 2 {
		for sum := 0; sum <= N; sum++ {
			if dp[pos][sum] == 0 {
				continue
			}
			for x := 1; x <= N; x++ {
				for y := 1; y <= N; y++ {
					prod := x * y
					if sum+prod <= N {
						dp[pos+2][sum+prod] += dp[pos][sum]
					}
				}
			}
		}
	}

	fmt.Println(dp[K][N])
}
