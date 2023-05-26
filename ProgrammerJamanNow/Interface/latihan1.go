package main

import (
	"fmt"
	"strconv"
)

type HasName interface {
	GetName()string
	GetUmur()int
	NameAndAge()string
}

func GetAll(hasName HasName) {
	fmt.Println(hasName.NameAndAge())
}



func SayHello(hasNamme HasName){
	fmt.Println("hello", hasNamme.GetName())
}

func UmurCalculate(hasName HasName)int{
	count := hasName.GetUmur() * 5
	return count
}

type Person struct{
	Name string
	Umur int
}

func (p Person) GetName() string{
	return p.Name
}

func (p Person) GetUmur() int{
	return p.Umur + 20
}

func (p Person) NameAndAge() string{
	data := strconv.Itoa(p.Umur)
	data = data + " "+ p.Name
	
	return  data
}



func main(){
	var people Person
	people.Name = "Eko"
	people.Umur = 20

	SayHello(people)
	GetAll(people)
	fmt.Println(UmurCalculate(people))
	
}