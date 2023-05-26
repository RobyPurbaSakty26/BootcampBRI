package main

import "fmt"

func main() {
	c := make(chan bool, 5)
	c <- true

	fmt.Println("this line runnig")

	data := <- c
	fmt.Println(data)
}