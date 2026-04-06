package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Digite o coeficiente A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o coeficiente B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o coeficiente C: ")
	fmt.Scan(&c)

	delta := b*b - 4*a*c

	fmt.Printf("Delta = %.2f\n", delta)

	if delta < 0 {
		fmt.Println("Classificação: RAÍZES IMAGINÁRIAS")
		fmt.Println("A equação não possui raízes reais.")

		parteReal := -b / (2 * a)
		parteImag := math.Sqrt(-delta) / (2 * a)

		fmt.Printf("X1 = %.4f + %.4fi\n", parteReal, parteImag)
		fmt.Printf("X2 = %.4f - %.4fi\n", parteReal, parteImag)

	} else if delta == 0 {
		fmt.Println("Classificação: RAIZ ÚNICA")

		x := -b / (2 * a)
		fmt.Printf("X1 = X2 = %.4f\n", x)

	} else {
		fmt.Println("Classificação: RAÍZES DISTINTAS")

		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)

		fmt.Printf("X1 = %.4f\n", x1)
		fmt.Printf("X2 = %.4f\n", x2)
	}

}
