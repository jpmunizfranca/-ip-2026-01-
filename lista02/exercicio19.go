package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64

	fmt.Print("Figura (1-Cone / 2-Cilindro / 3-Esfera): ")
	fmt.Scan(&opcao)

	if opcao == 1 {
		fmt.Print("Raio: ")
		fmt.Scan(&raio)
		fmt.Print("Altura: ")
		fmt.Scan(&altura)

		volume := (math.Pi * raio * raio * altura) / 3
		area := math.Pi * raio * math.Sqrt(raio*raio+altura*altura)

		fmt.Printf("Volume: %.4f\n", volume)
		fmt.Printf("Area: %.4f\n", area)

	} else if opcao == 2 {
		fmt.Print("Raio: ")
		fmt.Scan(&raio)
		fmt.Print("Altura: ")
		fmt.Scan(&altura)

		volume := math.Pi * raio * raio * altura
		area := 2 * math.Pi * raio * altura

		fmt.Printf("Volume: %.4f\n", volume)
		fmt.Printf("Area: %.4f\n", area)

	} else if opcao == 3 {
		fmt.Print("Raio: ")
		fmt.Scan(&raio)

		volume := (4.0 / 3.0) * math.Pi * raio * raio * raio
		area := 4 * math.Pi * raio * raio

		fmt.Printf("Volume: %.4f\n", volume)
		fmt.Printf("Area: %.4f\n", area)

	} else {
		fmt.Println("Opção inválida!")
	}
}
