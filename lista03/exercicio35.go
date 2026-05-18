package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite um número inteiro positivo na base 10: ")
	f.Scan(&n)

	binario := ""
	temp := n
	for temp > 0 {
		resto := temp % 2
		binario = f.Sprintf("%d", resto) + binario
		temp /= 2
	}
	if binario == "" {
		binario = "0"
	}
	f.Printf("%d em base 2: %s\n", n, binario)
}
