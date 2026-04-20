package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Print("Digite o primeiro número (N1): ")
	f.Scan(&n1)
	f.Print("Digite o segundo número (N2): ")
	f.Scan(&n2)

	resultado := 0
	for i := 0; i < n2; i++ {
		resultado += n1
	}
	f.Printf("%d x %d = %d\n", n1, n2, resultado)
}
