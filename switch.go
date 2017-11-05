package main

import "fmt"

func whatday(n int) {
	if n != 0 && n <= 7 {
		switch n {
		case 1:
			fmt.Println("Monday")
		default:
			fmt.Println("Not handled")
		}
	}
}
