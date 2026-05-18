package main

import "fmt"

func main() {
	var n int
	fmt.Println("Me informe o número de jogos.")
	fmt.Scan(&n)

	rendas := make([]float64, n)

	for i := 0; i < n; i++ {
		var p, g, a, c int
		fmt.Printf("Me informe a quantidade de pessoas (popular, geral, arquibancada, cadeiras) no jogo %d:\n", i+1)
		fmt.Scan(&p, &g, &a, &c)
		rendas[i] = float64(p*0 + g*5 + a*10 + c*20)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("A renda do jogo %d é = %.2f\n", i+1, rendas[i])
	}
}
