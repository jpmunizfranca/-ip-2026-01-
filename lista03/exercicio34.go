package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Print("Digite o primeiro número (N1): ")
	f.Scan(&n1)
	f.Print("Digite o segundo número (N2): ")
	f.Scan(&n2)

	mmc := n1
	for mmc%n2 != 0 {
		mmc += n1
	}
	f.Printf("MMC(%d, %d) = %d\n", n1, n2, mmc)
}
