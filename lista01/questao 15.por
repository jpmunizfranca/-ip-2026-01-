programa
{
    funcao inicio()
    {
        inteiro N, i, quadrado

        escreva("Digite o valor de N (5 < N < 2000): ")
        leia(N)
        
        para(i = 2; i <= N; i = i + 2)
        {
            quadrado = i * i

            escreva(i, "^2 = ", quadrado, "\n")
        }
    }
}
