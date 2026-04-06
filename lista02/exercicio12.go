package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	var classificacao string
	if idade <= 2 {
		classificacao = "Recém-nascido"
	} else if idade <= 11 {
		classificacao = "Criança"
	} else if idade <= 19 {
		classificacao = "Adolescente"
	} else if idade <= 55 {
		classificacao = "Adulto"
	} else {
		classificacao = "Idoso"
	}

	fmt.Printf("Idade: %d anos\nClassificação: %s\n", idade, classificacao)
}
