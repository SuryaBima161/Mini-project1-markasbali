package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DaftarBuku struct {
	Kode_Buku      string
	Judul_Buku     string
	Pengarang      string
	Penerbit       string
	Jumlah_Halaman string
	Tahun_Terbit   string
}

var listOrder []DaftarBuku

func TambahBuku() {
	judulbuku := ""
	pengarang := ""
	// penerbit := ""
	// jumlahhalaman := 0
	// tahunterbit := 0
	inputUser := bufio.NewReader(os.Stdin)
	fmt.Println("Tambah Kode Buku")
	fmt.Print("Silahkan Tambah List Buku :")
	// bufio.NewReader()
	kodebuku, err := inputUser.ReadString('\n')
	// _, err := fmt.Scanln(&order)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	kodebuku = strings.Replace(kodebuku, "\n", " ", 1)

	fmt.Println("Add table order")
	fmt.Print("Please add your table order :")

	_, err = fmt.Scanln(&judulbuku)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Add total order")
	fmt.Print("Please add your total order :")
	_, err = fmt.Scanln(&pengarang)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	listOrder = append(listOrder, DaftarBuku{
		Kode_Buku:  kodebuku,
		Judul_Buku: judulbuku,
		Pengarang:  pengarang,
	})
	fmt.Println("Order Successful ")
}

func listBuku() {
	// var buku string
}

func editBuku() {

}
func deleteBuku() {

}

func main() {

	fmt.Println("Hello Bang")
}
