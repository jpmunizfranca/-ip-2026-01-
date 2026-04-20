package main

import f "fmt"

func main() {
	var n int
	f.Println("Digite um valor n ( >= 0):")
	f.Scan(&n)
	if n <= 0 {
		f.Println("Programa encerrado.")
		return
	}
	fat := 1
	for i := 1; i <= n; i++ {
		fat *= i

	}
	f.Printf("O fatorial de %d é: %d", n, fat)
}
