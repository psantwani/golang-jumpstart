package main

import (
	"fmt"
	"time"
)

func goroutinedemo() {
	count(5)                    //simple execution
	go count(5)                 //ask it to run in the background
	time.Sleep(time.Second * 1) //delay of 1 second so that we can see the output before the program exits
}

func count(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}
