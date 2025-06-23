package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	groups := make([][]int, m)
	personToGroups := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var sz int
		fmt.Scan(&sz)
		group := make([]int, sz)
		for j := 0; j < sz; j++ {
			fmt.Scan(&group[j])
			personToGroups[group[j]] = append(personToGroups[group[j]], i)
		}
		groups[i] = group
	}

	fLevel := make([]int, n+1)
	visited := make([]bool, n+1)
	groupVisited := make([]bool, m)

	queue := []int{}
	for i := 0; i < k; i++ {
		var f0 int
		fmt.Scan(&f0)
		queue = append(queue, f0)
		visited[f0] = true
		fLevel[f0] = 0
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		level := fLevel[current]
		if level == 3 {
			continue
		}
		for _, groupIdx := range personToGroups[current] {
			if groupVisited[groupIdx] {
				continue
			}
			groupVisited[groupIdx] = true
			for _, member := range groups[groupIdx] {
				if !visited[member] {
					visited[member] = true
					fLevel[member] = level + 1
					queue = append(queue, member)
				}
			}
		}
	}

	cntF := [4]int{}
	for i := 1; i <= n; i++ {
		if visited[i] {
			cntF[fLevel[i]]++
		}
	}
	fmt.Println(cntF[1], cntF[2], cntF[3])
}
