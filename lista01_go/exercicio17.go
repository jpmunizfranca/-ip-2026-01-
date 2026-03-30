package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Me informe o valor de x e o valor de y.")
	fmt.Scan(&x, &y)
	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR.")
	} else {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x+i*2)
		}

	}
}
