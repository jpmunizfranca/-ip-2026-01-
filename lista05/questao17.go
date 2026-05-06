package main

import f "fmt"

func primo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var v [10]int
	f.Println("Digite 10 números inteiros:")
	for i := range v {
		f.Scan(&v[i])
	}

	encontrou := false
	for i, n := range v {
		if primo(n) {
			f.Printf("Número primo %d na posição %d\n", n, i)
			encontrou = true
		}
	}
	if !encontrou {
		f.Println("Nenhum número primo encontrado.")
	}
}
