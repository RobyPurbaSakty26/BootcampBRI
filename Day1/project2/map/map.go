package main

import (
	"fmt"
)




func main (){
	studentAges := map[string]int {
		"Jon": 20,
		"jane": 19,
	}
	studentAges["Roby"] = 23
	
	fmt.Println(studentAges["Jon"])
	fmt.Println(studentAges)
}
