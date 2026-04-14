package main

import "fmt"

func main() {
	var n int
	var x int

	fmt.Println("Me informe o tamanho da lista:")
	fmt.Scan(&n)

	l := make([]int, n)

	fmt.Println("Digite os elementos da lista:")
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	fmt.Println("Agora me informe um número x:")
	fmt.Scan(&x)

	resultado := lista(l, x)

	fmt.Println("Resultado:", resultado)
}

func lista(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
