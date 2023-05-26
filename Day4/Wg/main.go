package main

import (
	"fmt"
	"sync"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting \n and", id)

	fmt.Printf("Worker %d done \n", id)
}

func main(){
	var wg sync.WaitGroup

	for i:= 0 ; i <= 5; i ++{
		wg.Add(1)
		go Worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All worker have finish")
}