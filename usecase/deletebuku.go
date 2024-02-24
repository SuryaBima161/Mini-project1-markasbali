package usecase

import (
	"fmt"
)

func DeleteBuku() {
	var kode string
	fmt.Println("=================================")
	fmt.Println("Hapus Pesanan")
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
		}
	}
	if index != -1 {
		listBook = append(listBook[:index], listBook[index+1:]...)
		fmt.Println("Buku dengan Kode :", kode, " Berhasil di hapus")
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
}
