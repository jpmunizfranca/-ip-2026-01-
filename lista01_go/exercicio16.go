package main

import "fmt"

func main() {
	var salario, reajuste1, reajuste2 float64
	fmt.Println("Me informe o valor do salário.")
	fmt.Scan(&salario)
	if salario <= 300 {
		reajuste1 = 1.5 * salario
		fmt.Printf("Salário com reajuste é = %.2f", reajuste1)
	} else {
		reajuste2 = 1.3 * salario
		fmt.Printf("Salário com reajuste é = %2.f", reajuste2)
	}

}
