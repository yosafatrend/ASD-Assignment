package main

import "fmt"

func main() {
	var n, find int
	var data [10]int
	jml := 0

	fmt.Println("Mau input berapa data?")
	fmt.Scanln(&n)
	fmt.Println("Masukkan", n, "integer(s)")
	for i := 0; i < n; i++ {
		fmt.Scanln(&data[i])
	}

	fmt.Println("Cari angka berapa?")
	fmt.Scanln(&find)
	for i := 0; i < n; i++ {
		if data[i] == find {
			fmt.Println("Angka ", find, "KETEMU dan tersimpan di lokasi/indeks", i)
			jml++
		}
	}
	if jml == 0 {
		fmt.Println("Angka ", find, "TIDAK DITEMUKAN")
	}
}
