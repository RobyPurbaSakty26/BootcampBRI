package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		ch1 <- "Data dari chanel 1"
	}()

	go func ()  {
		time.Sleep(2 * time.Second)
		ch2 <- "Data dari chanel 2"
	}()

	select{
	case msg1 := <-ch1:
		fmt.Println("Menerima data dari channel 1: ", msg1)
	case msg2 := <- ch2:
		fmt.Println("Menerima data dari chanel 2: ", msg2)
	case <- time.After(3 * time.Second):
		fmt.Println("Time: Tidak ada data yang diterima dalam waktu yang ditentukan")
	}

}