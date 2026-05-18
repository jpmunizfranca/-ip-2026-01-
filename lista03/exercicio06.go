package main

import f "fmt"

func main() {
	var x int
	f.Println("Digite um número inteiro.")
	f.Scan(&x)
	if x <= 0 {
		return
	}
	encontrou := false
	for i := 1; i*(i+1)*(i+2) <= x; i++ {
		if i*(i+1)*(i+2) == x {
			encontrou = true
			break
		}
	}
	if encontrou {
		f.Println("Esse número é um número triangular.")
	} else {
		f.Println("Esse número não é um número triangular.")
	}

}
