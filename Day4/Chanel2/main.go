package main

import "fmt"

// membuat chanel dengan type data string
func main(){
	message := make(chan string)

	go func(){
		message <- "Hello, World"
	}()

	// menerima data dari channel
	receiveMsg := <- message

	fmt.Println(receiveMsg)
	
}