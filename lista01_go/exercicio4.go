package main

import "fmt"

func main() {
	var sal, qntd, custo3, custo1, custo2 float64

	fmt.Println("Me informe o valor do salário mínimo e a quantidade de kW.")
	fmt.Scan(&sal, &qntd)
	custo1 = (7 * sal) / 1000
	custo2 = custo1 * qntd
	custo3 = (custo2 * 9) / 10
	fmt.Println("Custo por kW: R$ ", custo1)
	fmt.Printf("Custo por consumo: R$%2.f", custo2)
	fmt.Printf("\nCusto com desconto: R$%2.f", custo3)

}
