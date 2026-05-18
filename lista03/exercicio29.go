package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite um número inteiro positivo (N): ")
	f.Scan(&n)

	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	f.Printf("Somatório de 1 a %d = %d\n", n, soma)
}
