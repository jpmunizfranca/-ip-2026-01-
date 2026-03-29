package main

import "fmt"

func main() {
	var a, b, c, media float64
	fmt.Println("Me fale as três notas de um aluno.")
	fmt.Scan(&a, &b, &c)
	media = (a + b + c) / 3
	fmt.Println("MEDIA = ", media)
	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
