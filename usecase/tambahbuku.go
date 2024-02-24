package usecase

import (
	"bufio"
	"fmt"
	"miniproject1/models"
	"os"
	"strings"
)

var listBook []models.DaftarBuku

func TambahBuku() {
	judulbuku := ""
	pengarang := ""
	penerbit := ""
	jumlahhalaman := 0
	tahunterbit := 0
	inputUser := bufio.NewReader(os.Stdin)
	fmt.Println("Tambah Kode Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	kodebuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	kodebuku = strings.Replace(kodebuku, "\n", " ", 1)

	fmt.Println("Tambah Judul Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	judulbuku, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	judulbuku = strings.Replace(judulbuku, "\n", " ", 1)

	fmt.Println("Tambah Pengarang Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	pengarang, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	pengarang = strings.Replace(pengarang, "\n", " ", 1)

	fmt.Println("Tambah nama penerbit")
	fmt.Print("Silahkan Tambah List Buku :")
	penerbit, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	penerbit = strings.Replace(penerbit, "\n", " ", 1)

	fmt.Println("Tambah Jumlah Halaman")
	fmt.Print("Silahkan Tambah List Buku :")
	_, err = fmt.Scanln(&jumlahhalaman)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Tambah tahun terbit")
	fmt.Print("Silahkan Tambah List Buku :")
	_, err = fmt.Scanln(&tahunterbit)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	listBook = append(listBook, models.DaftarBuku{
		Kode_Buku:      kodebuku,
		Judul_Buku:     judulbuku,
		Pengarang:      pengarang,
		Penerbit:       penerbit,
		Jumlah_Halaman: jumlahhalaman,
		Tahun_Terbit:   tahunterbit,
	})
	fmt.Println("Order Successful ")
}
