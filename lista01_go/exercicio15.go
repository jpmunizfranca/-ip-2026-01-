package main

import "fmt"

func main() {
	var x int
	fmt.Println("Me fale um valor inteiro maior que 5 e menor que 2000.")
	fmt.Scan(&x)
	for i := 2; i <= x; i += 2 {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}
