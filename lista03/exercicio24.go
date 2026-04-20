package main

import f "fmt"

func main() {
	f.Println("Tabela do seno usando a série de Mac-Laurin:")
	f.Printf("%-12s %-15s\n", "A (radianos)", "Sen(A)")

	for i := 0; i <= 63; i++ {
		a := float64(i) * 0.1
		sen := a - (a*a*a)/6.0 + (a*a*a*a*a)/120.0 - (a*a*a*a*a*a*a)/5040.0
		f.Printf("%-12.1f %-15.6f\n", a, sen)
	}
}
