package main

import (
	"fmt"
)

//Book struct
type Book struct {
	name   string
	author string
	pages  int
}

//Shelf struct
type Shelf struct {
	position int
	book     Book
}

func structdemo() {
	book := Book{
		name:   "Harry Potter",
		author: "J.K. Rowling",
		pages:  800,
	}

	s := Shelf{
		position: 1,
		book:     book,
	}

	fmt.Println(book)
	fmt.Println(s)
}
