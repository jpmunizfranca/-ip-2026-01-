package main

import "fmt"

func main() {
	var tempo, horas, minutos, segundos int
	fmt.Println("Me informe um valor de tempo em horas, minutos e segundos.")
	fmt.Scan(&horas, &minutos, &segundos)
	tempo = horas*3600 + minutos*60 + segundos
	fmt.Println("O tempo em segundos é = ", tempo)
}
