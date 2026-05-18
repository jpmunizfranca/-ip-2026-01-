package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Print("Digite o dividendo (N1): ")
	f.Scan(&n1)
	f.Print("Digite o divisor (N2): ")
	f.Scan(&n2)

	quociente := 0
	resto := n1
	for resto >= n2 {
		resto -= n2
		quociente++
	}
	f.Printf("Quociente(%d, %d) = %d\n", n1, n2, quociente)
	f.Printf("Resto(%d, %d) = %d\n", n1, n2, resto)
}
