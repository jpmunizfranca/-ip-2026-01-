package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Digite a idade:")
	fmt.Scan(&idade)
	if idade < 16 {
		fmt.Println("Tipo de eleitor: Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Tipo de eleitor: Eleitor obrigatório")
	} else {
		fmt.Println("Tipo de eleitor: Eleitor facultativo")
	}

}
