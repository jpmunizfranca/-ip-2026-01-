package main

import f "fmt"

func main() {
	var v [10]int
	f.Println("Digite 10 números inteiros:")
	for i := range v {
		f.Scan(&v[i])
	}

	contado := map[int]bool{}
	for i, n := range v {
		if contado[n] {
			continue
		}
		cont := 0
		for _, m := range v[i:] {
			if m == n {
				cont++
			}
		}
		if cont > 1 {
			f.Printf("Número %d aparece %d vezes\n", n, cont)
		}
		contado[n] = true
	}
}
