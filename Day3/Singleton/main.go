package main

import (
	"fmt"
	"sync"
)

// Singleton adalah struktur yang mewakili Singleton
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstance adalah metode untuk mendapatkan instance Singleton (hanya akan membuat instance pertama kali)
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Data dari Singleton"}
	})
	return instance
}

func main() {
	// Mendapatkan instance Singleton
	singleton := GetInstance()
	singleton1 := GetInstance()

	// Mengakses data Singleton
	fmt.Println(singleton.data, singleton1)
}
