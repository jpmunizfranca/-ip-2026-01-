package main

import f "fmt"

func main() {
	var b, n int
	f.Print("Digite a base b (b >= 2): ")
	f.Scan(&b)
	f.Print("Digite o expoente n (n > 1): ")
	f.Scan(&n)

	resultado := 1
	for i := 0; i < n; i++ {
		resultado *= b
	}
	f.Printf("%d^%d = %d\n", b, n, resultado)
}
