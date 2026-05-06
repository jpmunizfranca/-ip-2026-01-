package main

import f "fmt"

func main() {
	var alturas [10]float64
	f.Println("Digite a altura dos 10 atletas:")
	soma := 0.0
	for i := range alturas {
		f.Scan(&alturas[i])
		soma += alturas[i]
	}
	media := soma / 10

	f.Printf("Média: %.2f\n", media)
	f.Println("Atletas acima da média:")
	for i, h := range alturas {
		if h > media {
			f.Printf("Atleta %d: %.2f\n", i+1, h)
		}
	}
}
