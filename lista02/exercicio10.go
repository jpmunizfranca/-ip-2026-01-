package main

import "fmt"

func main() {
	precos := map[int][2]float64{
		1: {500.00, 900.00},
		2: {350.00, 650.00},
		3: {350.00, 600.00},
		4: {300.00, 550.00},
	}

	nomes := map[int]string{
		1: "Região Norte",
		2: "Região Nordeste",
		3: "Região Centro-Oeste",
		4: "Região Sul",
	}

	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")

	var destino int
	for {
		fmt.Print("Digite o número do destino: ")
		fmt.Scan(&destino)
		if destino >= 1 && destino <= 4 {
			break
		}
		fmt.Println("Opção inválida. Digite um número entre 1 e 4.")
	}

	fmt.Println("1 - Ida e Volta")
	fmt.Println("2 - Somente Ida")

	var tipo int
	for {
		fmt.Print("Digite o tipo da viagem (1 ou 2): ")
		fmt.Scan(&tipo)
		if tipo == 1 || tipo == 2 {
			break
		}
		fmt.Println("Opção inválida. Digite 1 ou 2.")
	}

	var preco float64
	var tipoStr string

	if tipo == 1 {
		preco = precos[destino][1]
		tipoStr = "Ida e Volta"
	} else {
		preco = precos[destino][0]
		tipoStr = "Somente Ida"
	}

	fmt.Printf("Destino: %s\n", nomes[destino])
	fmt.Printf("Tipo: %s\n", tipoStr)
	fmt.Printf("Preço: R$ %.2f\n", preco)
}
