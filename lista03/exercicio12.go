package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o valor de X: ")
	fmt.Scan(&x)

	soma := 0.0
	fat := 1.0
	sinal := 1.0

	for k := 0; k < 20; k++ {
		if k > 0 {
			fat = fat * float64(k)
		}
		soma = soma + sinal*(x/fat)
		sinal = sinal * -1
	}

	fmt.Printf("O resultado de S é : %.6f\n", soma)
}
