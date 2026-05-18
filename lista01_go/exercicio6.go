package main

import "fmt"

func main() {

	var a int
	fmt.Println("Me fale quantas temperaturas você quer transformar.")
	fmt.Scan(&a)
	temp := make([]float64, a)
	for i := 1; i <= a; i++ {
		fmt.Printf("Digite o valor da %d temperatura (em Fahrenheit):", i)
		fmt.Scan(&temp[i-1])
	}
	for i := 1; i <= a; i++ {
		celsius := (temp[i-1] - 32) * 5 / 9
		fmt.Printf("O valor da %dª temperatura é = %.2f°C\n", i, celsius)
	}
}
