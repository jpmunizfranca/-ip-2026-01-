package main

import (
	"fmt"
)

func main() {
	var conta int
	var tipo int
	var consumo float64
	var total float64

	fmt.Print("Digite o número da conta do cliente: ")
	fmt.Scan(&conta)

	fmt.Println("Digite 1 para consumidor residencial, 2 para comercial e 3 para indutrial.")
	fmt.Scan(&tipo)

	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	if tipo == 1 {
		total = 5 + (0.05 * consumo)
		fmt.Println("O número da conta é =", conta)
		fmt.Printf("O consumo é = %.2f", total)
	}
	if tipo == 2 {
		total = 500 + 0.25*(consumo-80)
		fmt.Println("O número da conta é =", conta)
		fmt.Printf("O consumo é = %.2f", total)
	}
	if tipo == 3 {
		total = 800 + 0.04*(consumo-100)
		fmt.Println("O número da conta é =", conta)
		fmt.Printf("O consumo é = %.2f", total)
	}
}
