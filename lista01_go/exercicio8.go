package main

import "fmt"

func main() {
	var r, h, custo float64
	var area, ab, al float64
	fmt.Println("Me fale o valor do raio da base do cilindro.")
	fmt.Scan(&r)
	fmt.Println("Me fale o valor da altura do cilindro.")
	fmt.Scan(&h)
	ab = 3.14159 * r * r
	al = 2 * 3.14159 * r * h
	area = 2*ab + al
	custo = 100 * area
	fmt.Printf("O valor do custo é = %.2f", custo)

}
