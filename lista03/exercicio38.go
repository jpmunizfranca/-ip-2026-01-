package main

import f "fmt"

func main() {
	var cpf [11]int

	for i := 0; i < 11; i++ {
		f.Scan(&cpf[i])
	}

	soma := 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma += cpf[i] * peso
		peso--
	}

	resto := soma % 11
	d1 := 0
	if resto >= 2 {
		d1 = 11 - resto
	}

	soma = 0
	peso = 11

	for i := 0; i < 10; i++ {
		soma += cpf[i] * peso
		peso--
	}

	resto = soma % 11
	d2 := 0
	if resto >= 2 {
		d2 = 11 - resto
	}

	if d1 == cpf[9] && d2 == cpf[10] {
		f.Println("Valido")
	} else {
		f.Println("Invalido")
	}
}
