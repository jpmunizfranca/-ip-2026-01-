package main

import f "fmt"

func binario(n int) {
	if n == 0 {
		return
	}
	binario(n / 2)
	f.Print(n % 2)
}

func main() {
	var numero int
	f.Print("Digite um numero decimal: ")
	f.Scan(&numero)

	f.Print("O valor em binario é: ")
	if numero == 0 {
		f.Print(0)
	} else {
		binario(numero)
	}
	f.Println()
}
