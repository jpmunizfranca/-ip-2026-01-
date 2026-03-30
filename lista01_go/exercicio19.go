package main

import "fmt"

func main() {
	var soma float64
	var n int
	fmt.Println("Me infome o valor de n.")
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		soma += 1 / float64(i)
	}
	fmt.Printf("Soma aproximada = %.6f", soma+1)
}
