package main

import (
	"fmt"
	"miniproject1/usecase"
	"os"
)

type DaftarBuku struct {
	Kode_Buku      string
	Judul_Buku     string
	Pengarang      string
	Penerbit       string
	Jumlah_Halaman int
	Tahun_Terbit   int
}

func listBuku() {
	// var buku string
}

func editBuku() {

}
func deleteBuku() {

}

func main() {

	chooseOrder := 0
	fmt.Println("Order Food System")
	fmt.Println("================================================")
	fmt.Println("choose your order :")
	fmt.Println("1. tambah buku")
	fmt.Println("2. view order")
	fmt.Println("3. delete order")
	fmt.Println("4. Out System")
	fmt.Println("Enter your order")
	_, err := fmt.Scanln(&chooseOrder)

	if err != nil {
		fmt.Println("error: ", err)
	}

	if chooseOrder == 1 {
		usecase.TambahBuku()
	} else if chooseOrder == 2 {
		editBuku()
	} else if chooseOrder == 3 {
		deleteBuku()
	} else if chooseOrder == 4 {
		listBuku()
	} else if chooseOrder == 5 {
		os.Exit(0)
	}
	main()

	fmt.Println("Hello Bang")
}
