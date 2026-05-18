package main

import f "fmt"

func main() {
	soma := 0
	media := 0
	for i := 50; i <= 70; i += 2 {
		soma = soma + i
		media = soma / 11
	}
	f.Println("A soma é:", soma)
	f.Println("A média é:", media)
}
