package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	sort.Ints(cp)

	m := make(map[int]int)
	for _, n := range cp {
		if _, ok := m[n]; !ok {
			m[n] = len(m) + 1
		}
	}

	for i, n := range arr {
		arr[i] = m[n]
	}
	return arr
}

func main() {
	case1 := []int{40, 10, 20, 40, 30}
	case2 := []int{100, 100, 100}
	case3 := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(case1)) // [4, 1, 2, 4, 3]
	fmt.Println(arrayRankTransform(case2)) // [1, 1, 1]
	fmt.Println(arrayRankTransform(case3)) // [5,3,4,2,8,6,7,1,3]
}
