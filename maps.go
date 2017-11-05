package main

import (
	"fmt"
)

func mapdemo() {
	nm := map[string]int{"first": 1, "second": 2}
	fmt.Println("Map nm: ", nm)

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2

	fmt.Println("Map m: ", m)

	A := m["a"]
	fmt.Println("Value of A: ", A)

	delete(m, "b")
	fmt.Println("Deleted b from m: ", m)

	B, isBpresent := m["b"]
	A, isApresent := m["a"]

	fmt.Println("Value of B is: ", B)
	fmt.Println("Value of A is: ", A)
	fmt.Println("isBpresent", isBpresent)
	fmt.Println("isApresent", isApresent)

}
