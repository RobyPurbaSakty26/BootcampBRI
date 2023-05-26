package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"John" : 30,
		"Alice" : 25,
		"Bob" : 35,
	}

	// mengases nulai dalam map menggunakan kunci
	fmt.Println((ages["John"]))
	fmt.Println((ages["Alice"]))
}