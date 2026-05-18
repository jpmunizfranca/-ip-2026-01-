package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo de 3 casas: ")
	fmt.Scan(&numero)

	if numero >= 100 && numero <= 999 {
		dezena := (numero / 10) % 10
		fmt.Println("Algarismo da casa das dezenas:", dezena)
	} else {
		fmt.Println("O número informado não possui 3 casas.")
	}

}
