package main

import "fmt"

func main() {
	var a1, r, n int
	fmt.Println("Me informe o valor do primeiro termo da p.a., sua razão e sua quantidade de termos. ")
	fmt.Scan(&a1, &r, &n)

	soma := 0
	termo := a1

	for i := 0; i < n; i++ {
		soma += termo
		termo += r
	}

	fmt.Println(soma)
}
