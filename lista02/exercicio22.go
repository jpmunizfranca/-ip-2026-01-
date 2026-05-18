package main

import "fmt"

func main() {
	var matricula, horasExtras int

	fmt.Print("Matrícula do funcionário: ")
	fmt.Scan(&matricula)
	fmt.Print("Horas extras trabalhadas: ")
	fmt.Scan(&horasExtras)

	salarioMinimo := 788.00
	valorHoraExtra := 10.00

	salarioHoraExtra := float64(horasExtras) * valorHoraExtra
	salarioBruto := 3*salarioMinimo + salarioHoraExtra

	inss := 0.0
	ir := 0.0

	if salarioBruto > 1500.00 {
		inss = salarioBruto * 0.12
	}
	if salarioBruto > 2000.00 {
		ir = salarioBruto * 0.20
	}

	salarioLiquido := salarioBruto - inss - ir

	fmt.Printf("\nMatrícula: %d\n", matricula)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", inss)
	fmt.Printf("Desconto IR:   R$ %.2f\n", ir)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}
