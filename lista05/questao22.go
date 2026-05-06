package main

import f "fmt"

func encontrarConta(codigos [10]int, cod int) int {
	for i, c := range codigos {
		if c == cod {
			return i
		}
	}
	return -1
}

func main() {
	var codigos [10]int
	var saldos [10]float64

	f.Println("Cadastre 10 contas (código e saldo):")
	for i := 0; i < 10; i++ {
		f.Printf("Conta %d - Código: ", i+1)
		f.Scan(&codigos[i])
		f.Printf("Conta %d - Saldo: ", i+1)
		f.Scan(&saldos[i])
	}

	for {
		f.Println("\n1. Depósito\n2. Saque\n3. Ativo bancário\n4. Sair")
		var op int
		f.Print("Opção: ")
		f.Scan(&op)

		switch op {
		case 1:
			var cod int
			var valor float64
			f.Print("Código da conta: ")
			f.Scan(&cod)
			idx := encontrarConta(codigos, cod)
			if idx == -1 {
				f.Println("Conta não encontrada.")
				continue
			}
			f.Print("Valor a depositar: ")
			f.Scan(&valor)
			saldos[idx] += valor
			f.Printf("Novo saldo: %.2f\n", saldos[idx])

		case 2:
			var cod int
			var valor float64
			f.Print("Código da conta: ")
			f.Scan(&cod)
			idx := encontrarConta(codigos, cod)
			if idx == -1 {
				f.Println("Conta não encontrada.")
				continue
			}
			f.Print("Valor a sacar: ")
			f.Scan(&valor)
			if valor > saldos[idx] {
				f.Println("Saldo insuficiente.")
			} else {
				saldos[idx] -= valor
				f.Printf("Saque realizado. Novo saldo: %.2f\n", saldos[idx])
			}

		case 3:
			total := 0.0
			for _, s := range saldos {
				total += s
			}
			f.Printf("Ativo bancário total: %.2f\n", total)

		case 4:
			f.Println("Encerrando...")
			return
		}
	}
}
