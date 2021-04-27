package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	s2 := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)

	// sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sort.IntSlice(s2).Sort()
	fmt.Println(s2)
}
