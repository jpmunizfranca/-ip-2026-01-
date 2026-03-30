package main

import "fmt"

func main() {
	var consumo, conta, total1, total2, total3 float64
	var x string
	fmt.Println("Me informe o número de sua conta, seu consumo e seu tipo de consumidor: residencial(r), comercial(c), ou industrial(i).")
	fmt.Scan(&conta, &consumo, &x)
	if x == "r" {
		total1 = 5 + 0.05*consumo
		fmt.Println("CONTA = ", conta)
		fmt.Println("VALOR DA CONTA = R$", total1)
	}
	if x == "c" {
		total2 = 500 + 0.25*(consumo-80)
		fmt.Println("CONTA = ", conta)
		fmt.Println("VALOR DA CONTA = R$", total2)
	}
	if x == "i" {
		total3 = 800 + 0.04*(consumo-100)
		fmt.Println("CONTA =", conta)
		fmt.Println("VALOR DA CONTA = R$", total3)
	}
}
