package main

import f "fmt"

func main() {
	f.Println("Índices de todos os elementos da matriz 10x10:")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			f.Printf("(%d,%d) ", i, j)
		}
		f.Println()
	}
}
