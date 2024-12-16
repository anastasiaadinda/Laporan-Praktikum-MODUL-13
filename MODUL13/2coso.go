package main

import "fmt"

func main() {
	var number int
	var countinuLoop bool
	for countinuLoop = true; countinuLoop; {
		fmt.Scan(&number)
		countinuLoop = number <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", number)
}
