package main

import (
	f "fmt"
	"sort"
)

type Empregado struct {
	codigo, meses int
}

func main() {
	var emp []Empregado
	for {
		var cod, meses int
		f.Print("Código e meses (0 0 para sair): ")
		f.Scan(&cod, &meses)
		if cod == 0 && meses == 0 {
			break
		}
		emp = append(emp, Empregado{cod, meses})
	}

	sort.Slice(emp, func(i, j int) bool {
		return emp[i].meses < emp[j].meses
	})

	f.Println("Os 3 empregados mais recentes:")
	for i := 0; i < 3 && i < len(emp); i++ {
		f.Printf("Código: %d, Meses: %d\n", emp[i].codigo, emp[i].meses)
	}
}
