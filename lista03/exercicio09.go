package main

import f "fmt"

func main() {
	var n int
	var nota1, nota2 float64
	var media, mediat, mediatotal float64
	f.Println("Informe a quantidade de alunos:")
	f.Scan(&n)

	reprovado := 0
	exame := 0
	aprovado := 0

	for i := 1; i <= n; i++ {
		f.Printf("\nDigite a nota 1 do %.d aluno: ", i)
		f.Scan(&nota1)
		f.Printf("Digite a nota 2 do %.d aluno: ", i)
		f.Scan(&nota2)
		media = (float64(nota1 + nota2)) / 2
		if media <= 3 {
			reprovado = reprovado + 1
		}
		if media > 3 && media < 7 {
			exame = exame + 1
		}
		if media >= 7 {
			aprovado = aprovado + 1
		}

		mediat += media
		mediatotal = float64(mediat) / float64(n)
		f.Printf("A média do %.d aluno é: %.2f\n", i, media)

	}
	f.Println("\nRESULTADOS FINAIS:")
	f.Printf("O total de alunos aprovados: %d\n", aprovado)
	f.Printf("O total de alunos reprovados: %d\n", reprovado)
	f.Printf("O total de alunos de exame: %d\n", exame)
	f.Printf("A média da sala: %.2f\n", mediatotal)

}
