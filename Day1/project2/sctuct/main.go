package main

import (
	"fmt"
)

// type Person struct {
// 	Name string
// 	Age int
// }

// type Departement struct {
// 	Name string
// 	Manager *Person
// }

func main () {
	// dept := Departement{
	// 	Name: "Sales",
	// }
	// dept.Manager.Name = "Jhoe Doe"
	// dept.Manager.Age = 40

	// fmt.Println(dept.Manager.Name, dept.Manager.Age)

for i := 1; i <= 4; i++ {

    for j := 1; j <= i; j++ {

        fmt.Print("*")

    }

    fmt.Println()

}

for i := 3; i >= 1; i-- {

    for j := 1; j <= i; j++ {

        fmt.Print("*")

    }

    fmt.Println()

}

}