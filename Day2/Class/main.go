package main

import "fmt"

// Manusia class
type Manusia struct {
	Nama   string
	Usia   int
	Gender string
}

// GetNama method mengembalikan nama manusia
func (m *Manusia) GetNama() string {
	return m.Nama
}

// GetUsia method mengembalikan usia manusia
func (m *Manusia) GetUsia() int {
	return m.Usia
}

// GetGender method mengembalikan jenis kelamin manusia
func (m *Manusia) GetGender() string {
	return m.Gender
}

// SetNama method mengatur nama manusia
func (m *Manusia) SetNama(nama string) {
	m.Nama = nama
}

// SetUsia method mengatur usia manusia
func (m *Manusia) SetUsia(usia int) {
	m.Usia = usia
}

// SetGender method mengatur jenis kelamin manusia
func (m *Manusia) SetGender(gender string) {
	m.Gender = gender
}

func main() {
	// Membuat instance dari kelas Manusia
	manusia := &Manusia{Nama: "John Doe", Usia: 30, Gender: "Laki-laki"}

	// Mengakses dan mengubah atribut
	fmt.Println("Nama Manusia:", manusia.GetNama())
	fmt.Println("Usia Manusia:", manusia.GetUsia())
	fmt.Println("Jenis Kelamin Manusia:", manusia.GetGender())

	manusia.SetNama("Jane Doe")
	manusia.SetUsia(25)
	manusia.SetGender("Perempuan")

	fmt.Println("Nama Manusia setelah perubahan:", manusia.GetNama())
	fmt.Println("Usia Manusia setelah perubahan:", manusia.GetUsia())
	fmt.Println("Jenis Kelamin Manusia setelah perubahan:", manusia.GetGender())
}
