package main

import "fmt"

func main() {
	var notas = [3]float64{}
	var media float64
	fmt.Println("Me fale a primeira nota.")
	fmt.Scan(&notas[0])
	fmt.Println("Me fale a segunda nota.")
	fmt.Scan(&notas[1])
	fmt.Println("Me fale a terceira nota.")
	fmt.Scan(&notas[2])
	for i := 0; i < 3; i++ {
		media = (notas[0] + notas[1] + notas[2]) / 3
		fmt.Printf("A %d nota é %.1f \n", i+1, notas[i])
	}
	fmt.Printf("A média é = %.2f \n", media)
	switch {
	case notas[0] >= media:
		fmt.Printf("A %.1f nota está dentro da média.\n", notas[0])

	case notas[0] < media:
		fmt.Printf("A %.1f nota está abaixo da média.\n", notas[0])
	}
	switch {
	case notas[1] >= media:
		fmt.Printf("A %.1f nota está dentro da média.\n", notas[1])

	case notas[1] < media:
		fmt.Printf("A %.1f nota está abaixo da média.\n", notas[1])
	}
	switch {
	case notas[2] >= media:
		fmt.Printf("A %.1f nota está dentro da média.\n", notas[2])

	case notas[2] < media:
		fmt.Printf("A %.1f nota está abaixo da média.\n", notas[2])
	}

}
