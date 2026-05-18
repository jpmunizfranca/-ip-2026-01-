package main

import "fmt"

func main() {
	var a, b, c int
	var x, y int
	fmt.Println("Me fale três números.")
	fmt.Scan(&a, &b, &c)

	if a >= 10 || b >= 10 || c >= 10 || a < 1 || b < 1 || c < 1 {
		fmt.Println("DÍGITO INVÁLIDO")
	} else {
		x = 100*a + 10*b + c
		y = x * x
		fmt.Println(x, " , ", y)
	}

}
