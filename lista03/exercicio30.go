package main

import (
	f "fmt"
	"math"
)

func main() {
	f.Println("Tabela de volumes de esferas:")
	f.Printf("%-12s %-15s\n", "Raio (cm)", "Volume (cm³)")

	for i := 0; ; i++ {
		r := float64(i) * 0.5
		if r > 20 {
			break
		}
		volume := (4.0 / 3.0) * math.Pi * r * r * r
		f.Printf("%-12.1f %-15.6f\n", r, volume)
	}
}
