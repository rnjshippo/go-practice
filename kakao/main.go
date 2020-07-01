package main

import (
	"fmt"
	"sort"
)

// ByCost 커스텀 정렬
type ByCost [][]int

func (b ByCost) Len() int {
	return len(b)
}
func (b ByCost) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByCost) Less(i, j int) bool {
	return b[i][2] < b[j][2]
}

var parent [105]int

func find(a int) int {
	if a == parent[a] {
		return a
	}
	parent[a] = find(parent[a])
	return parent[a]
}

func merge(a, b int) {
	parent[parent[a]] = parent[b]
}

func solution(n int, costs [][]int) int {
	var ret int = 0
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	sort.Sort(ByCost(costs))
	for _, num := range costs {
		if find(num[0]) != find(num[1]) {
			merge(num[0], num[1])
			ret += num[2]
		}
	}
	return ret
}
func main() {
	var test = [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}}
	fmt.Println(solution(5, test))
}
