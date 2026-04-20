package main

import f "fmt"

func main() {
	total := uint64(0)
	graos := uint64(1)

	for i := 1; i <= 64; i++ {
		total += graos
		if i < 64 {
			graos *= 2
		}
	}
	f.Printf("Total de grãos no tabuleiro de xadrez: %d\n", total)
}
