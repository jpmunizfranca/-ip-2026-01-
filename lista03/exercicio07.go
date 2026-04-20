package main

import f "fmt"

func main() {
	var x int
	soma := 0
	somap := 0
	qntdp := 0
	qntdi := 0
	i := 1

	var maior, menor int
	primeiro := true

	for ; i < 100; i++ {
		f.Printf("Digite o %d número: ", i)
		f.Scan(&x)

		if x == 30000 {
			break
		}

		soma += x

		// maior e menor
		if primeiro {
			maior = x
			menor = x
			primeiro = false
		} else {
			if x > maior {
				maior = x
			}
			if x < menor {
				menor = x
			}
		}

		if x%2 == 0 {
			somap += x
			qntdp++
		} else {
			qntdi++
		}
	}

	total := qntdp + qntdi

	var media float64
	var mediap float64
	var porcentagemi float64

	if total > 0 {
		media = float64(soma) / float64(total)
		porcentagemi = float64(qntdi) / float64(total) * 100
	}

	if qntdp > 0 {
		mediap = float64(somap) / float64(qntdp)
	}

	f.Printf("\nSoma = %d", soma)
	f.Printf("\nQuantidade = %d", total)
	f.Printf("\nMédia = %.2f", media)
	f.Printf("\nMaior número = %d", maior)
	f.Printf("\nMenor número = %d", menor)
	f.Printf("\nMédia dos pares = %.2f", mediap)
	f.Printf("\nPorcentagem de ímpares = %.2f%%", porcentagemi)
}
