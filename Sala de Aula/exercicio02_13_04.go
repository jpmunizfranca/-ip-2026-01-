package main

import "fmt"

func main() {
	l := []int{10, 20, 30, 40, 50, 60, 70, 80}

	var x int
	fmt.Println("Me fale um número:")
	fmt.Scan(&x)

	resultado := lista(l, x)

	fmt.Println("Resultado:", resultado)
}

func lista(l []int, x int) int {
	e := 0
	d := len(l) - 1

	for e <= d {
		m := (e + d) / 2

		if l[m] == x {
			return m
		}

		if l[m] < x {
			e = m + 1
		} else {
			d = m - 1
		}
	}

	return -1
}
