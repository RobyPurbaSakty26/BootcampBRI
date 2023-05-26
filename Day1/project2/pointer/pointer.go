package main

import (
	"fmt"
)

func IncrementByTwo (n *int){
	*n += 2
}

func main () {
	num := 5 
	IncrementByTwo(&num)
	fmt.Println(num)
}