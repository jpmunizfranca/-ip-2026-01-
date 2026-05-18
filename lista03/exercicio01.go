package main

import f "fmt"

func main() {
	var x, y int
	f.Println("Me fale o valor da base x.")
	f.Scan(&x)
	f.Println("Me fale o valor do expoente y.")
	f.Scan(&y)
	resultado := 1
	for i := 0; i < y; i++ {
		resultado = resultado * x
	}
	f.Println("O valor do resultado é ", resultado)
}
