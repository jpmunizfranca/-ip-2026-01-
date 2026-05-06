package main

import (
	f "fmt"
	"math"
)

func main() {
	var b [100]float64
	f.Println("Digite 100 valores:")
	for i := range b {
		f.Scan(&b[i])
	}

	S := 0.0
	for i := 0; i < 50; i++ {
		diff := b[i] - b[99-i]
		S += math.Pow(diff, 3)
	}
	f.Printf("Somatório S = %.4f\n", S)
}
