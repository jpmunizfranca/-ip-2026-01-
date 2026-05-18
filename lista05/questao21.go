package main

import f "fmt"

func main() {
	var v [10]float64
	f.Println("Digite os 10 números reais do vetor:")
	for i := range v {
		f.Scan(&v[i])
	}

	for {
		var cod int
		f.Print("Digite o código (0=sair, 1=direto, 2=inverso): ")
		f.Scan(&cod)
		switch cod {
		case 0:
			return
		case 1:
			f.Println("Ordem direta:", v)
		case 2:
			for i := len(v) - 1; i >= 0; i-- {
				f.Printf("%.2f ", v[i])
			}
			f.Println()
		default:
			f.Println("Código inválido.")
		}
	}
}
