package main

import f "fmt"

func main() {
	s := 0.0
	for k := 1; k <= 37; k++ {
		a := float64(39 - k)
		b := float64(38 - k)
		s += a * b / float64(k)
	}
	f.Printf("S = %.6f\n", s)
}
