package main

import (
	"fmt"
)

func main () {
	numbers := []int{1,2,3,4,5}

	// mengases element slice menggunakan index
	fmt.Println("nilai idex 0 ",numbers[0])
	fmt.Println("nilai idex 2 ", numbers[2])

	// mengubah nilai slice
	numbers[1] = 10
	fmt.Println(numbers)


	// menambahkan element slice
	numbers = append(numbers, 6)
	fmt.Println(numbers)
}