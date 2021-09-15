package main

import "fmt"

func main() {
	var bangunDatar hitung
	var alas, tinggi, sisi float64
	var choice int
	var isExit bool
	var done string

	for isExit == false {
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======Menu=======")
		fmt.Println("1. Luas segitiga")
		fmt.Println("2. Keliling segitiga")
		fmt.Println("3. Luas persegi")
		fmt.Println("4. Keliling persegi")
		fmt.Println("0. Keluar")
		fmt.Println("=================")

		fmt.Print("Masukan pilihan : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Print("Masukan alas : ")
				fmt.Scan(&alas)
				fmt.Print("Masukan tinggi : ")
				fmt.Scan(&tinggi)
				bangunDatar = segitiga{alas: alas, tinggi: tinggi}
				fmt.Println("Luas segitiga :", bangunDatar.luas())
				fmt.Print("Masukan string untuk kembali : ")
				fmt.Scan(&done)
			}
		case 2:
			{
				fmt.Print("Masukan alas : ")
				fmt.Scan(&alas)
				bangunDatar = segitiga{alas: alas}
				fmt.Println("Keliling segitiga :", bangunDatar.keliling())
				fmt.Print("Masukan string untuk kembali : ")
				fmt.Scan(&done)
			}
		case 3:
			{
				fmt.Print("Masukan sisi : ")
				fmt.Scan(&sisi)
				bangunDatar = persegi{sisi: sisi}
				fmt.Println("Luas Persegi :", bangunDatar.luas())
				fmt.Print("Masukan string untuk kembali : ")
				fmt.Scan(&done)
			}
		case 4:
			{
				fmt.Print("Masukan sisi : ")
				fmt.Scan(&sisi)
				bangunDatar = persegi{sisi: sisi}
				fmt.Println("Keliling Persegi :", bangunDatar.keliling())
				fmt.Print("Masukan string untuk kembali : ")
				fmt.Scan(&done)
			}
		case 0:
			{
				isExit = true
			}
		}
	}
}

type hitung interface {
	luas() float64
	keliling() float64
}

type segitiga struct {
	alas   float64
	tinggi float64
}

func (s segitiga) luas() float64 {
	return 0.5 * s.alas * s.tinggi
}

func (s segitiga) keliling() float64 {
	return 3 * s.alas
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}