package main

import "fmt"

func main() {
	var x int
	fmt.Println("Me fale um número.")
	fmt.Scan(&x)
	if x%3 == 0 {
		fmt.Println("Esse número é divisível por 3.")
	}
	if x%5 == 0 {
		fmt.Println("Esse número é divisível por 5.")
	}
	if x%3 != 0 && x%5 != 0 {
		fmt.Println("Esse número não é divisível nem por 5 e nem por 3.")
	}
}
