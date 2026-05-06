package main

import f "fmt"

func potencia(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}

func main() {
	var x, n int
	f.Print("Digite o numero base: ")
	f.Scan(&x)
	f.Print("Digite o expoente: ")
	f.Scan(&n)
	f.Println("O valor final é:", potencia(x, n))
}
