programa
{
    funcao inicio()
    {
        real nota
        cadeia conceito

        
        escreva("Digite a nota (0 a 10): ")
        leia(nota)

        
        se(nota >= 9.0 e nota <= 10.0)
        {
            conceito = "A"
        }
        senao se(nota >= 7.5 e nota < 9.0)
        {
            conceito = "B"
        }
        senao se(nota >= 6.0 e nota < 7.5)
        {
            conceito = "C"
        }
        senao se(nota >= 0.0 e nota < 6.0)
        {
            conceito = "D"
        }
        senao
        {
            conceito = "INVALIDA"
        }

      
        escreva("NOTA = ", nota, " CONCEITO = ", conceito, "\n")
    }
}
