package main

import (
	"fmt"
)

func main (){

	// ini sialisasi awal
	var numbers [5]int
	
	 numbers = [5]int{1,2,3,4,5}

	 var isArray [10]int
	 fmt.Println(isArray)

	//  akses element array menggunakan indexs
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])

	// mengubag nilai element array
	numbers[1] = 10

	fmt.Println(numbers)



}