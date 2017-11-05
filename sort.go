package main

import (
	"fmt"
	"sort"
)

func sortdemo() {
	s := []string{"z", "x", "a", "q", "v"}
	sort.Strings(s)
	fmt.Println("Sorted strings: ", s)

	nums := []int{10, 4, 8, 101, 2}
	sort.Ints(nums)
	fmt.Println("Sorted numbers: ", nums)
}
