package main

import "fmt"

func pointer1() {
	var a *int
	var b = 6

	a = &b
	fmt.Println("Address: ", a)
	fmt.Println("Value: ", *a)
}

func pointer2() {
	c := 5
	increment(c)
	increment(c)
	fmt.Println("Called from pointer2 func - part 1 : ", c)

	incrementbyptr(&c)
	incrementbyptr(&c)
	fmt.Println("Called from pointer2 func - part 2 : ", c)
}

func increment(c int) {
	c++
	fmt.Println("Called within increment", c)
}

func incrementbyptr(c *int) {
	*c++
	fmt.Println("Called within incrementbyptr", *c)
}
