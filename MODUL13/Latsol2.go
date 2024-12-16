package main

import (
	"fmt"
	"math"
)

func main() {

	var masukanDesimal float64

	fmt.Print("Massukan bilangan desimal: ")
	fmt.Scan(&masukanDesimal)

	//membuat ke atas angka input
	//ceil -> untuk menentukan bilangan bulat ke atas dari input
	atas := math.Ceil(masukanDesimal)

	//inisialisasi nilai awal
	nilai := masukanDesimal

	//loop hinggga mencapai bilangan bulat (1)
	for nilai < atas {
		nilai = math.Round((nilai+0.1)*10) / 10
		fmt.Printf("%.1f\n", nilai)
	}

}
