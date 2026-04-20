package main

import f "fmt"

func main() {
	var carlos float64
	f.Println("Escreva o valor do salário de Carlos.")
	f.Scan(&carlos)
	joao := carlos / 3
	rendacarlos := carlos
	rendajoao := joao
	for i := 1; i <= 100; i++ {
		rendacarlos = rendacarlos * 1.02
		rendajoao = rendajoao * 1.05
		if rendajoao >= rendacarlos {
			f.Println("O montante de João é maior ou igual ao de Carlos após", i, "meses.")
			f.Printf("Após %.d meses, o montante de João é ~= %.2f e o de Carlos é ~= %.2f", i, rendajoao, rendacarlos)
			break
		}
	}
}
