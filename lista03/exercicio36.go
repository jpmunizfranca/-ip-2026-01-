package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite um número inteiro positivo na base 10: ")
	f.Scan(&n)

	hex := ""
	digitos := "0123456789ABCDEF"
	temp := n
	for temp > 0 {
		resto := temp % 16
		hex = string(digitos[resto]) + hex
		temp /= 16
	}
	if hex == "" {
		hex = "0"
	}
	f.Printf("%d em base 16: %s\n", n, hex)
}
