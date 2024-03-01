package usecase

import (
	"fmt"
)

var kode string

func DeleteBuku() {

	fmt.Println("=================================")
	fmt.Println("Delete Buku")
	fmt.Println("=================================")
	ListBuku()
	fmt.Println("=================================")
	fmt.Print("Masukan Kode Buku : ")
	_, err := fmt.Scanln(&kode)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
	index := -1
	for i, buku := range listBook {
		if buku.Kode_Buku == kode {
			index = i
			fmt.Println("Buku dengan Kode : ", buku.Kode_Buku, " dengan judul : ", buku.Judul_Buku)
			break
		}
	}
	if index != -1 {
		listBook = append(listBook[:index], listBook[index+1:]...)
		fmt.Println("Berhasil Dihapus")
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
}
