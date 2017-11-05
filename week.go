package main

import "github.com/grsmv/goweek"
import "fmt"

func days() {
	week, _ := goweek.NewWeek(2017, 46)
	fmt.Println(week.Days)
}
