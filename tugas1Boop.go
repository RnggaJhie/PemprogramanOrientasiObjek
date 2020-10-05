package main

import "fmt"

type murid struct{
    nama string;
    nourut int;
    jurusan string;
}

func main(){
	var mrd murid
	fmt.Println("MEMASUKAN INPUTAN NAMA MURID")

	for i := 0; i < 6; i++{
	
	fmt.Println("Masukan Nama :")
	fmt.Scan(&mrd.nama)
	fmt.Println("Masukan Nourut :")
	fmt.Scan(&mrd.nourut)
	fmt.Println("Masukan Jurusan :")
	fmt.Scan(&mrd.jurusan)
	}

}
