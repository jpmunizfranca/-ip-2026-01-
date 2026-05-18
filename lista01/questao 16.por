programa
{
    inclua biblioteca Matematica --> mat

    funcao inicio()
    {
        real salario, sr

        escreva("Digite o salário do funcionário.")
        leia(salario)

        se(salario <= 300.00){
        sr = salario + (salario * 0.50)
        }
        senao{
        sr = salario + (salario * 0.30)
        }

        sr = mat.arredondar(sr, 2)

        escreva("SALARIO COM REAJUSTE = ", sr, "\n")
    }
}
