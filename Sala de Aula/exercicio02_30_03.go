package main

import "fmt"

func main() {
	var n = [5]int{}
	var soma int
	fmt.Println("Insira o primeiro valor:")
	fmt.Scan(&n[0])
	fmt.Println("Insira o segundo valor:")
	fmt.Scan(&n[1])
	fmt.Println("Insira o terceiro valor:")
	fmt.Scan(&n[2])
	fmt.Println("Insira o quarto valor:")
	fmt.Scan(&n[3])
	fmt.Println("Insira o quinto valor:")
	fmt.Scan(&n[4])

	for i := 0; i < 5; i++ {
		soma = n[0] + n[1] + n[2] + n[3] + n[4]
		fmt.Printf("O %d valor é = %d\n ", i+1, n[i])
	}
	fmt.Println("A soma é igual a ", soma)
}
