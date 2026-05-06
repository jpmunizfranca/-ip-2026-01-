package main

import f "fmt"

func inverte(arr []int, n int) {
	if n <= 1 {
		return
	}
	arr[0], arr[n-1] = arr[n-1], arr[0]
	inverte(arr[1:n-1], n-2)
}

func main() {
	var tamanho int
	f.Print("Qual o tamanho do array? ")
	f.Scan(&tamanho)

	numeros := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		f.Print("Digite um numero inteiro: ")
		f.Scan(&numeros[i])
	}

	inverte(numeros, tamanho)
	f.Println("O array invertido é:", numeros)
}
