package usecase

import (
	"bufio"
	"fmt"
	"miniproject1/models"
	"os"
	"strings"
)

var listBook []models.DaftarBuku
var inputUser = bufio.NewReader(os.Stdin)

func TambahBuku() {
	jumlahhalaman := 0
	tahunterbit := 0

	fmt.Println("Tambah Kode Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	kodebuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	kodebuku = strings.TrimSpace(kodebuku)

	fmt.Println("Tambah Judul Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	judulbuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	judulbuku = strings.TrimSpace(judulbuku)

	fmt.Println("Tambah Pengarang Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	pengarang, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	pengarang = strings.TrimSpace(pengarang)

	fmt.Println("Tambah nama penerbit")
	fmt.Print("Silahkan Tambah List Buku :")
	penerbit, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	penerbit = strings.TrimSpace(penerbit)

	fmt.Println("Tambah Jumlah Halaman")
	fmt.Print("Silahkan Tambah List Buku :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &jumlahhalaman)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Tambah tahun terbit")
	fmt.Print("Silahkan Tambah List Buku :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &tahunterbit)
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
	fmt.Println("Tambah Buku Berhasil ")
}
