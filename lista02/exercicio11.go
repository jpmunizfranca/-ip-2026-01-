package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("Indefinido: divisão por zero.")
	} else {
		resultado := 8.0 / (2.0 - x)
		fmt.Printf("f(%.2f) = %.2f\n", x, resultado)
	}
}
