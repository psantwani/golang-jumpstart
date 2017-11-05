package main

import "fmt"

func arr1() {
	var a [7]int
	a[0] = 1
	a[6] = 7

	fmt.Println(a)
}

func arr2() {
	a := [3]int{1, 2, 3}
	fmt.Println(len(a))
}

func matrix() {
	var m [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = 1
		}
	}

	fmt.Println(m)
}
