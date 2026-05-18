package main

import "fmt"

func main() {
	var preco float64
	var codigo int

	fmt.Print("Preço do produto: R$ ")
	fmt.Scan(&preco)
	fmt.Print("Condição de pagamento (1-4): ")
	fmt.Scan(&codigo)

	if codigo == 1 {
		total := preco * 0.90
		fmt.Printf("À vista (dinheiro/cheque) com 10%% de desconto: R$ %.2f\n", total)
	} else if codigo == 2 {
		total := preco * 0.95
		fmt.Printf("À vista (cartão) com 5%% de desconto: R$ %.2f\n", total)
	} else if codigo == 3 {
		parcela := preco / 2
		fmt.Printf("2x sem juros: R$ %.2f cada parcela\n", parcela)
	} else if codigo == 4 {
		total := preco * 1.10
		parcela := total / 3
		fmt.Printf("3x com 10%% de juros: R$ %.2f cada parcela (total: R$ %.2f)\n", parcela, total)
	} else {
		fmt.Println("Código inválido!")
	}
}
