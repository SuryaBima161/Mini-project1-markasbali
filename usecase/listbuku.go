package usecase

import "fmt"

func ListBuku() {
	fmt.Println("List Book")
	for index, list := range listBook {
		fmt.Printf("%d. Kode Buku : %s Judul Buku : %s Pengarang %s : Penerbit %s : Jumlah Halaman %d : Tahun Terbit %d\n",
			index+1, list.Kode_Buku, list.Judul_Buku, list.Pengarang, list.Penerbit, list.Jumlah_Halaman, list.Tahun_Terbit) //kenapa index+1 karena pada sebuah slice dimuali dari 0, maka +1 sudah dimulai dari
	}
}
