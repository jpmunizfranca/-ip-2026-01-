package main

import "fmt"

func main() {
	var dia, categoria int
	var preco float64

	fmt.Print("Dia da semana (1-Domingo, 2-Seg, 3-Ter, 4-Qua, 5-Qui, 6-Sex, 7-Sab): ")
	fmt.Scan(&dia)
	fmt.Print("Categoria (1-Comum, 2-Lancamento): ")
	fmt.Scan(&categoria)
	fmt.Print("Preco normal do DVD: ")
	fmt.Scan(&preco)

	if dia == 2 || dia == 3 || dia == 5 {
		preco = preco * 0.60
	}
	if categoria == 2 {
		preco = preco * 1.15
	}

	fmt.Printf("Preco final: R$ %.2f\n", preco)
}
