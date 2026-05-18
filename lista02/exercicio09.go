package main

import "fmt"

func main() {
	var compra, venda float64
	fmt.Println("Fale o valor da compra do comerciante.")
	fmt.Scan(&compra)
	if compra < 10 {
		venda = 1.7 * compra
		fmt.Printf("O valor da venda é %.2f", venda)
	}
	if compra >= 10 && compra < 30 {
		venda = 1.5 * compra
		fmt.Printf("O valor da venda é %.2f", venda)
	}
	if compra >= 30 && compra < 50 {
		venda = 1.4 * compra
		fmt.Printf("O valor da venda é %.2f", venda)
	}
	if compra >= 50 {
		venda = 1.3 * compra
		fmt.Printf("O valor da venda é %.2f", venda)
	}
}
