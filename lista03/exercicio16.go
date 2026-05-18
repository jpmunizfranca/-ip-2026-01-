package main

import f "fmt"

func main() {
	var n, a1, a2 int
	f.Print("Digite o primeiro termo da série de Fetuccine: ")
	f.Scan(&a1)
	f.Print("Digite o segundo termo da série de Fetuccine: ")
	f.Scan(&a2)
	f.Print("Digite o número de termos (N >= 3): ")
	f.Scan(&n)

	f.Printf("Série de Fetuccine: %d, %d", a1, a2)

	prev2, prev1 := a1, a2
	for i := 3; i <= n; i++ {
		var atual int
		if i%2 != 0 {
			atual = prev1 + prev2
		} else {
			atual = prev1 - prev2
		}
		f.Printf(", %d", atual)
		prev2 = prev1
		prev1 = atual
	}
}
