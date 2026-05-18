package main

import f "fmt"

func mostrarDisponiveis(poltronas [24]int, tipo string) {
	f.Printf("Poltronas disponíveis na %s: ", tipo)
	alguma := false
	for i, v := range poltronas {
		if v == 0 {
			f.Printf("%d ", i+1)
			alguma = true
		}
	}
	if !alguma {
		f.Printf("nenhuma")
	}
	f.Println()
}

func main() {
	var janela, corredor [24]int

	for {
		totalLivre := 0
		for _, v := range janela {
			if v == 0 {
				totalLivre++
			}
		}
		for _, v := range corredor {
			if v == 0 {
				totalLivre++
			}
		}
		if totalLivre == 0 {
			f.Println("Ônibus completamente lotado!")
			return
		}

		var tipo int
		f.Print("\nDeseja poltrona: 1=Janela, 2=Corredor, 0=Sair: ")
		f.Scan(&tipo)
		if tipo == 0 {
			return
		}

		var poltronas *[24]int
		nome := ""
		if tipo == 1 {
			poltronas = &janela
			nome = "janela"
		} else {
			poltronas = &corredor
			nome = "corredor"
		}

		livres := 0
		for _, v := range poltronas {
			if v == 0 {
				livres++
			}
		}
		if livres == 0 {
			f.Printf("Não há poltronas livres no %s.\n", nome)
			continue
		}

		mostrarDisponiveis(*poltronas, nome)
		var poltrona int
		f.Print("Escolha o número da poltrona: ")
		f.Scan(&poltrona)
		if poltrona < 1 || poltrona > 24 || poltronas[poltrona-1] == 1 {
			f.Println("Poltrona inválida ou já ocupada.")
		} else {
			poltronas[poltrona-1] = 1
			f.Printf("Poltrona %d no %s reservada com sucesso!\n", poltrona, nome)
		}
	}
}
