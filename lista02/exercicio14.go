package main

import "fmt"

func main() {
	var preco float64
	var opcao int

	fmt.Print("Digite o preço inicial de fábrica: R$ ")
	fmt.Scan(&preco)

	fmt.Println("Opcionais disponíveis:")
	fmt.Println("1 - Ar condicionado     R$ 1750,00")
	fmt.Println("2 - Pintura metálica    R$  800,00")
	fmt.Println("3 - Vidro elétrico      R$ 1200,00")
	fmt.Println("4 - Direção hidráulica  R$ 2000,00")

	adicionais := map[int]struct {
		nome  string
		valor float64
	}{
		1: {"Ar condicionado", 1750.00},
		2: {"Pintura metálica", 800.00},
		3: {"Vidro elétrico", 1200.00},
		4: {"Direção hidráulica", 2000.00},
	}

	selecionados := map[int]bool{}

	for {
		fmt.Print("Digite o número do opcional desejado (0 para finalizar): ")
		fmt.Scan(&opcao)

		if opcao == 0 {
			break
		}

		if _, existe := adicionais[opcao]; !existe {
			fmt.Println("Opcional inválido.")
			continue
		}

		if selecionados[opcao] {
			fmt.Printf("%s já foi adicionado.\n", adicionais[opcao].nome)
			continue
		}

		selecionados[opcao] = true
		preco += adicionais[opcao].valor
		fmt.Printf("%s adicionado.\n", adicionais[opcao].nome)
	}

	fmt.Printf("\nPreço final do carro: R$ %.2f\n", preco)
}
