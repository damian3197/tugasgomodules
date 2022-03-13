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
func Dendabuku(hargaDenda, jumlahhari int) int{
	denda := hargaDenda * jumlahhari
	return denda
	}

// struct
type DataMember struct {
	Nama string
	PinjamBuku bool
	StatusMember bool
	JudulBuku string
	LamaPeminjaman int
}
func DataPengunjung(DataMember) {
	datamember := DataMember {
		Nama: "Damian Sitanggang",
		StatusMember: false,
		PinjamBuku: true,
		JudulBuku: "Buku Sains",
		LamaPeminjaman: 5,
	}
	fmt.Println(datamember)
}
// struct Method
func (datamember DataMember) CekBuku(buku string) {
	fmt.Println("Anda meminjam Buku berjudul ", datamember.JudulBuku)
}

// function multiple return value
func PowerCal(I int, V int) (int, int) {
	Power := V * I
	Resistor := V % I
	return Power, Resistor
}

// Anonymous Struct
func AnonymousDataBuku() {
	DataBuku := struct {
	Buku string
	Judul string
	StatusBuku bool
	}{
		Buku: "Novel",
		Judul: "Olympus",
		StatusBuku: true,
	}
	fmt.Println(DataBuku)	
}


