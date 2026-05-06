package main

import f "fmt"

func main() {
	var v [10]int
	for i := range v {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&v[i])
	}

	encontrou := false
	for i, n := range v {
		if n > 50 {
			f.Printf("Número %d na posição %d\n", n, i)
			encontrou = true
		}
	}
	if !encontrou {
		f.Println("Nenhum número é superior a 50.")
	}
}
