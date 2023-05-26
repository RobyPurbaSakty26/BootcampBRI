package main

import (
	"fmt"
	"time"
)

func taksA() {
	for i := 0; i < 5; i++ {
		fmt.Println("Taks A : ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func taksB() {
	for i := 0; i < 5; i++ {
		fmt.Println("Taks B : ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func taksC(c chan<- string){
	time.Sleep(time.Second * 3)
	c <- "Taks C Dijalankan"
}

func taksD(c chan<- string){
	time.Sleep(time.Second * 3)
	c <- "Taks D Dijalankan"
}

func main() {
	// go taksA()
	// go taksB()

	// time.Sleep(time.Second * 3)

	ch := make(chan string)

	go taksC(ch)
	go taksD(ch)

	x, y := <-ch, <-ch
	fmt.Println(x)
	fmt.Println(y)
}