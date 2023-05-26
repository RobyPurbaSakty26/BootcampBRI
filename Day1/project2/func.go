package main

import "fmt"

// nill adalag nilai kosong di dalam go
type Increment func(int) int
func IncrementByOne (val int)int {
	return val+1
}

func AppllyOpration(name string, val int, opration Increment) int {
	fmt.Println("Opration name ", name)
	return opration(val)
}



func main() {
	result := AppllyOpration("increment", 10, IncrementByOne)
	fmt.Println(result)
}