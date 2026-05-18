package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, me float64

	fmt.Print("Número do aluno: ")
	fmt.Scan(&id)
	fmt.Print("Nota 1: ")
	fmt.Scan(&n1)
	fmt.Print("Nota 2: ")
	fmt.Scan(&n2)
	fmt.Print("Nota 3: ")
	fmt.Scan(&n3)
	fmt.Print("Média dos exercícios: ")
	fmt.Scan(&me)

	media := (n1 + n2*2 + n3*3 + me) / 7

	conceito := ""
	if media >= 9.0 {
		conceito = "A"
	} else if media >= 7.5 {
		conceito = "B"
	} else if media >= 6.0 {
		conceito = "C"
	} else if media >= 4.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	situacao := "REPROVADO"
	if conceito == "A" || conceito == "B" || conceito == "C" {
		situacao = "APROVADO"
	}

	fmt.Printf("\nAluno: %d\n", id)
	fmt.Printf("Notas: %.1f / %.1f / %.1f\n", n1, n2, n3)
	fmt.Printf("Média dos exercícios: %.1f\n", me)
	fmt.Printf("Média final: %.2f\n", media)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Situação: %s\n", situacao)
}
