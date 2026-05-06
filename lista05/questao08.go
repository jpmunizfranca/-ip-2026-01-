package main

import (
	f "fmt"
	"math"
)

func main() {
	var v [15]float64
	f.Println("Digite 15 números inteiros:")
	for i := range v {
		var n float64
		f.Scan(&n)
		if n < 0 {
			v[i] = -1
		} else {
			v[i] = math.Sqrt(n)
		}
	}
	f.Println("Raízes quadradas:", v)
}
