package main

import "fmt"

func main() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(i int) {
			// can do anyting
			fmt.Println(i ,"Hello")
		}(v)
	}
}