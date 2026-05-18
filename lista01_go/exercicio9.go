package main

import "fmt"

func main() {
	var a, b, c, delta float64
	fmt.Println("Me fale o valor de a:")
	fmt.Scan(&a)
	fmt.Println("Me fale o valor de b:")
	fmt.Scan(&b)
	fmt.Println("Me fale o valor de c:")
	fmt.Scan(&c)
	delta = b*b - 4*a*c
	fmt.Printf("O valor de delta é = %.2f", delta)
}
