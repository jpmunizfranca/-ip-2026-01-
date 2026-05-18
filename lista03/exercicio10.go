package main

import f "fmt"

func main() {
	var n int
	f.Println("Digite um número n ( >= 3) do tamanho da sequência de Fibonacci.")
	f.Scan(&n)
	a, b := 0, 1
	f.Println(a)
	f.Println(b)
	for i := 2; i < n; i++ {
		c := a + b
		f.Println(c)
		a = b
		b = c
	}
}
