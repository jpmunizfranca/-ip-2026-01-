package main

import f "fmt"

func main() {
	var idMaisGordo, idMaisMagro int
	var pesoMaisGordo, pesoMaisMagro float64

	for i := 1; i <= 90; i++ {
		var id int
		var peso float64
		f.Printf("Boi %d - Digite o número de identificação: ", i)
		f.Scan(&id)
		f.Printf("Boi %d - Digite o peso (kg): ", i)
		f.Scan(&peso)

		if i == 1 {
			idMaisGordo = id
			pesoMaisGordo = peso
			idMaisMagro = id
			pesoMaisMagro = peso
		} else {
			if peso > pesoMaisGordo {
				pesoMaisGordo = peso
				idMaisGordo = id
			}
			if peso < pesoMaisMagro {
				pesoMaisMagro = peso
				idMaisMagro = id
			}
		}
	}

	f.Printf("\nBoi mais gordo: ID %d | Peso: %.2f kg\n", idMaisGordo, pesoMaisGordo)
	f.Printf("Boi mais magro: ID %d | Peso: %.2f kg\n", idMaisMagro, pesoMaisMagro)
}
