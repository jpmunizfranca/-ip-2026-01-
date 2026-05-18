package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Println("Digite N1 e N2:")
	f.Scan(&n1, &n2)

	for i := n1; i <= n2; i++ {
		cont := 0

		for j := 1; j <= i; j++ {
			if i%j == 0 {
				cont++
			}
		}

		if cont == 2 {
			f.Println(i)
		}
	}
}
