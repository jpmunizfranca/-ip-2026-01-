package main

import f "fmt"

func main() {
	f.Println("Índices da diagonal principal da matriz 10x10:")
	for i := 1; i <= 10; i++ {
		f.Printf("(%d,%d)\n", i, i)
	}
}
