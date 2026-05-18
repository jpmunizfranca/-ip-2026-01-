package main

import "fmt"

func main() {
    var idade,opcao, count50, countAltura, countPeso, totalPessoas int
    var altura, peso, somaAltura float64
    
    for {
        fmt.Println("Digite a idade:")
        fmt.Scan(&idade)

        fmt.Println("Digite a altura:")
        fmt.Scan(&altura)

        fmt.Println("Digite o peso:")
        fmt.Scan(&peso)

        totalPessoas++

        if idade > 50 {
            count50++
        }

        if idade >= 10 && idade <= 20 {
            somaAltura += altura
            countAltura++
        }

        if peso < 40 {
            countPeso++
        }

        fmt.Println("Deseja continuar? 1 - Sim | outro valor - Não")
        fmt.Scan(&opcao)

        if opcao != 1 {
            break
        }
    }
