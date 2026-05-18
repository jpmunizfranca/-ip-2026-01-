package main

import (
	"fmt"
)

func main() {
	meses := [13]string{
		"",
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	var dia, mes, ano int
	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	fmt.Print("Digite o mês (número): ")
	fmt.Scan(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	if mes < 1 || mes > 12 {
		fmt.Println("Mês inválido! Digite um valor entre 1 e 12.")
		return
	}
	fmt.Printf("%d de %s de %d\n", dia, meses[mes], ano)
}
