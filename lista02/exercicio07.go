package main

import "fmt"

func main() {
	var A, B, C int

	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&A)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&B)
	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&C)

	var MENOR, INTER, MAIOR int

	if A < B && A < C {
		MENOR = A
	} else if B < A && B < C {
		MENOR = B
	} else {
		MENOR = C
	}

	if A > B && A > C {
		MAIOR = A
	} else if B > A && B > C {
		MAIOR = B
	} else {
		MAIOR = C
	}

	if A != MENOR && A != MAIOR {
		INTER = A
	} else if B != MENOR && B != MAIOR {
		INTER = B
	} else {
		INTER = C
	}

	fmt.Printf("\nOrdem crescente: %d %d %d\n", MENOR, INTER, MAIOR)
}
