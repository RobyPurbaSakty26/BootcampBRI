package main

import (
	"fmt"
)

func PrintOddMain(ch chan int) {
	for i := 0; i <= 5; i++ {
		fmt.Println("\t func 1 : ", i)
	}
	ch <- 0
}
func PrintEvenMain(ch chan int) {
	for i := 6; i <= 10; i++ {
		fmt.Println("func 2 : ", i)
	}
	ch <- 1
}

func main (){

	ch := make(chan int)

	go PrintOddMain(ch)
	go PrintEvenMain(ch)

	
	fmt.Println("Finish", <- ch)
	fmt.Println("Finish", <- ch)
	
	// time.Sleep(20 * time.Second)
}