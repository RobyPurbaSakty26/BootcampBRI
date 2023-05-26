package pointer

import (
	"fmt"
)

func Pointer(){
	num := 42
	var ptr *int

	ptr = &num

	fmt.Println("ini num ", num)
	fmt.Println("ini ptr", ptr)

	*ptr = 20

	fmt.Println("ini num ", num)
	fmt.Println("ini ptr", ptr)
}