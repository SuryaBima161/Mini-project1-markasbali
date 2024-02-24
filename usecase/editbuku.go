package usecase

import (
	"fmt"
	"strings"
)

func EditBuku() {
	fmt.Println("=================================")
	fmt.Println("Edit Buku")
	fmt.Println("=================================")
	ListBuku()
	fmt.Println("=================================")
	fmt.Print("Masukan Kode Buku : ")
	var kode string
	_, err := fmt.Scanln(&kode)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}
	index := -1
	for i, buku := range listBook {
		if buku.Kode_Buku == kode {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Buku tidak ditemukan")
		return
	}

	var judulBaru, pengarangBaru, penerbitBaru string
	var jumlahHalamanBaru, tahunterbitBaru int

	fmt.Print("Judul Baru: ")
	judulBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	judulBaru = strings.TrimSpace(judulBaru)

	fmt.Print("Pengarang Baru: ")
	pengarangBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	pengarangBaru = strings.TrimSpace(pengarangBaru)

	fmt.Print("Penerbit Baru: ")
	penerbitBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	penerbitBaru = strings.TrimSpace(penerbitBaru)

	fmt.Print("Jumlah Halaman Baru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &jumlahHalamanBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("Tahun Terbit Buru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &tahunterbitBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	listBook[index].Judul_Buku = judulBaru
	listBook[index].Pengarang = pengarangBaru
	listBook[index].Penerbit = penerbitBaru
	listBook[index].Jumlah_Halaman = jumlahHalamanBaru
	listBook[index].Tahun_Terbit = tahunterbitBaru

	fmt.Println("Data buku berhasil diperbarui.")
}
