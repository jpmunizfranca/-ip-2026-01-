package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite um número inteiro positivo na base 8: ")
	f.Scan(&n)

	decimal := 0
	base := 1
	temp := n
	for temp > 0 {
		digito := temp % 10
		decimal += digito * base
		base *= 8
		temp /= 10
	}
	f.Printf("%d (base 8) = %d (base 10)\n", n, decimal)
}
