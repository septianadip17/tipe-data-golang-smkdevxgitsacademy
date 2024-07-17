package main

import "fmt"

func main() {
	// // string
	// var name string
	// name = "Helmi"
	// name2 := "Adi"
	// var name3 = "Prasetyo"
	// fmt.Println(name, name2, name3)

	// // integer
	// var x, y int
	// x = 5
	// y = 7
	// result := x * y

	// // float64
	// var v, w float64
	// v = 2
	// w = 4
	// resulttwo := v / w
	// fmt.Println(result)
	// fmt.Println(resulttwo)

	// conditional
	var angka1, angka2 int64
	angka1 = 99
	angka2 = 2

	resultIf := angka1 % angka2

	var isGenap bool
	if resultIf == 0 {
		isGenap = true
	} else {
		isGenap = false
	}

	if isGenap {
		fmt.Println("Dia Bilangan Genap")
	} else {
		fmt.Println("Dia Bilangan Gajil")
	}
}
