package main

import "fmt"

func slice1() {
	x := make([]int, 0) //length is 0
	x = append(x, 1, 2, 3, 4, 5)

	fmt.Println(x)
}

func slice2() {
	x := make([]int, 3, 10) //start with 3 zeros and capacity of 10 (which can be modified)

	x[1] = 1 // 0 1 0

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)

	fmt.Println("x->", x)

	fmt.Println("x[2:5]", x[2:5])

	fmt.Println("x[:5]", x[:5])

	fmt.Println("x[:]", x[:])

}

func slice3() {
	s := make([]string, 3) // return [   ]
	s[0] = "a"
	s = append(s, "p", "q")
	fmt.Println(s)
}
