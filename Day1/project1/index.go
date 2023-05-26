// string golang di buat menggunakan rune
// rune => type data untuk string
// bila menggunakan huruf kecil maka variable tersebut akan hanya dapat digunakan pada package itu sendiri

package main

import (
	"test/pointer"
)

type Person struct {
	Name string 
}



func main () {
	// fmt.Println("Hello")
	pointer.Pointer()
}