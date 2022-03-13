package modul

import "fmt"

// function
func Welcome() {
	fmt.Println("Perpustakaan Kota")
}

// function sebagai parameter
func Verifikasi(namaDepan string, namaBelakang string) {
	fmt.Println("Verifikasi", namaDepan, namaBelakang, "Selamat Datang Di Perpustakaan")
}

// funtion dengan return value
func Dendabuku(hargaDenda, LamaPeminjaman int) int{
	denda := hargaDenda * LamaPeminjaman
	return denda
	}