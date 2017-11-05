package main

import "fmt"

func loop1() {
	for i := 0; i <= 7; i++ {
		fmt.Println(i)
		if i == 4 {
			break
		}
	}
}
