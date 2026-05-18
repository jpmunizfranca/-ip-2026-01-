package main

import f "fmt"

func main() {
	f.Println("Tabela de lucros esperados por preço de ingresso:")
	f.Printf("%s %s %s\n", "Preço (R$)", "Ingressos", "Lucro (R$)")

	preco := 6.0
	ingressos := 130
	despesas := 300.0
	lucroMax := -1e18
	precoMax := 0.0
	ingressosMax := 0

	for preco >= 1.0 {
		receita := preco * float64(ingressos)
		lucro := receita - despesas
		f.Printf("%-16.2f %-14d %-12.2f\n", preco, ingressos, lucro)

		if lucro > lucroMax {
			lucroMax = lucro
			precoMax = preco
			ingressosMax = ingressos
		}

		preco -= 0.60
		ingressos += 30
	}

	f.Printf("\nLucro máximo esperado: R$ %.2f\n", lucroMax)
	f.Printf("Preço correspondente:  R$ %.2f\n", precoMax)
	f.Printf("Ingressos correspondentes: %d\n", ingressosMax)
}
