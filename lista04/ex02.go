package main

import f "fmt"

func soma(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + soma(arr[1:])
}

func main() {
	var tamanho int
	f.Print("Quantos numeros reais voce quer somar? ")
	f.Scan(&tamanho)

	valores := make([]float64, tamanho)
	for i := 0; i < tamanho; i++ {
		f.Print("Digite um numero real: ")
		f.Scan(&valores[i])
	}

	f.Println("A soma total é:", soma(valores))
}
