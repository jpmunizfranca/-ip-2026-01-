package main

import "fmt"

func main() {
	var a, b, c, d, det float64
	fmt.Println("Me fale o valor de a:")
	fmt.Scan(&a)
	fmt.Println("Me fale o valor de b:")
	fmt.Scan(&b)
	fmt.Println("Me fale o valor de c:")
	fmt.Scan(&c)
	fmt.Println("Me fale o valor de d:")
	fmt.Scan(&d)
	det = a*d - b*c
	fmt.Printf("O valor do determinante é = %.2f", det)
}
