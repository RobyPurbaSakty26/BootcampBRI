package main

import (
	"fmt"
)

func main (){
	var number int = 10

	var pointer *int = &number

	fmt.Println(pointer)
	fmt.Println(number)

	*pointer = 50

	fmt.Println(number)
	// mengambil value dari pointer
	fmt.Println(*pointer)
}