package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
"Roby",
  23,
	}

	fmt.Println(p.Age)
}