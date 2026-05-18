package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Digite um número inteiro (<= 0 para sair): ")
		fmt.Scan(&n)

		if n <= 0 {
			break
		}

		quadrado := false

		for i := 1; i*i <= n; i++ {
			if i*i == n {
				quadrado = true
				break
			}
		}

		if quadrado {
			fmt.Println(n, "é um quadrado perfeito.")
		} else {
			fmt.Println(n, "não é um quadrado perfeito.")
		}
	}
}
