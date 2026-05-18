programa
{
    funcao inicio()
    {
        inteiro x, y, i, contador
        escreva("Escreva dois números inteiros.")
        leia(x, y)

        se(x % 2 == 0){
            contador = 0
            i = x

            enquanto(contador < y){
            
                escreva(i, " ")
                i = i + 2 
                contador = contador + 1}
            escreva("\n")
        }
        senao{
        escreva("O PRIMEIRO NUMERO NAO E PAR\n")}
        
    }
}
