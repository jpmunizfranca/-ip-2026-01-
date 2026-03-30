package main

import "fmt"

func main() {
	var total1, total2, total3 float64
	var horas int
	fmt.Println("Me fale o total de horas locadas.")
	fmt.Scan(&horas)
	total1 = float64(horas/3) * 10
	total2 = float64(horas % 3)
	total3 = total1 + total2*5
	fmt.Printf("O valor total a pagar é = %.2f", total3)
}
