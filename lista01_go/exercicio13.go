package main

import "fmt"

func main() {
	var nota float64
	fmt.Println("Me fale uma nota de 0 a 10 e com no máximo uma casa decimal.")
	fmt.Scan(&nota)
	if 9 <= nota && nota <= 10 {
		fmt.Println("Nota = ", nota)
		fmt.Println("Conceito = A")
	}
	if 7.5 <= nota && nota < 9 {
		fmt.Println("Nota= ", nota)
		fmt.Println("Conceito = B")
	}
	if 6 <= nota && nota < 7.5 {
		fmt.Println("Nota = ", nota)
		fmt.Println("Conceito = C")
	}
	if 0 <= nota && nota < 6 {
		fmt.Println("Nota = ", nota)
		fmt.Println("Conceito = D")
	}
}
