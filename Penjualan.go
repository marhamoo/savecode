package main

import "fmt"

func main () {

	//Deklarasi variabel
	var Nama_Customer, Nama_Barang string 
	var Harga float32 = 25000
	var Quantity int32
	var Hasil_Discount, Total_Harga float32

	//Input Nama Customer 
	fmt.Print("Input Nama Customer: ")
	fmt.Scanf("%s\n", &Nama_Customer)

	// Input Nama Barang
	fmt.Print("Input Nama Barang: ")
	fmt.Scanf("%s\n", &Nama_Barang)

	// Input Quantity Barang
	fmt.Print("Input Quantity Barang: ")
	fmt.Scanf("%d\n", &Quantity)

	// Kondisi Diskon, kalau lebih dari 5 dapat %10, selain itu %2 
	if Quantity > 5 { 
		Hasil_Discount = 0.1 // 10%
	} else {
		Hasil_Discount = 0.02 // 2%
	}

	// Proses 
	sub_total := float32(Quantity) * Harga
	Total_Harga = sub_total - (Hasil_Discount * sub_total)

	// Tampilkan Hasil Harga
	fmt.Println("Hasil_Discount: ", Hasil_Discount)
	fmt.Println("Quantity: ", Quantity)
	fmt.Println("Harga: ", Harga)
	fmt.Println("Total_Harga: ", Total_Harga)
}