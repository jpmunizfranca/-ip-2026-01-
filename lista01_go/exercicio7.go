package main

import "fmt"

func main() {
	var far, celsius, pol, mil float64
	fmt.Println("Me informe o valor de graus Fahrenheit.")
	fmt.Scan(&far)
	fmt.Println("Me informe a quantidade da chuva em polegadas.")
	fmt.Scan(&pol)
	celsius = (5*far - 160) / 9
	mil = pol / 25.4
	fmt.Printf("O valor em Celsius é = %.2f ", celsius)
	fmt.Printf("O quantidade de chuva em mm é = %.2f ", mil)
}
