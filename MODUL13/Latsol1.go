package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&input)

	if input <= 0 {
		fmt.Println("Masukan harus berupa bilangan bulat positif.")
		return
	}

	count := 0
	for input > 0 {
		input /= 10 
		count++     
	}

	fmt.Print(count)
}
