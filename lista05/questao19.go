package main

import f "fmt"

func main() {
	var nums [10]int
	var divs [5]int

	f.Println("Digite 10 números:")
	for i := range nums {
		f.Scan(&nums[i])
	}
	f.Println("Digite 5 divisores:")
	for i := range divs {
		f.Scan(&divs[i])
	}

	for _, n := range nums {
		f.Printf("Número %d:\n", n)
		for j, d := range divs {
			if d != 0 && n%d == 0 {
				f.Printf("  Divisível por %d na posição %d\n", d, j)
			}
		}
	}
}
