package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, v, ab float64
	raiz := math.Sqrt(3)
	fmt.Println("Me fale o valor da altura e da aresta de uma pirâmidade de base hexagonal.")
	fmt.Scan(&h, &a)
	ab = (3 * a * a * raiz) / 2
	v = (ab * h) / 3
	fmt.Printf("O volume da pirâmide é = %.2f metros cúbicos", v)
}
