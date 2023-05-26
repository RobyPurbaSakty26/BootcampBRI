package main

import "fmt"

type Shape interface {
	Area() int
}

type Kotak struct {
	sisi int
}

func (k Kotak) Area() int {
	return k.sisi * k.sisi
}

func ShowArea(s Shape) {
	fmt.Print(s.Area())

}

func main (){
	ktk := Kotak{10}

	ShowArea(ktk)
}