programa
{
    funcao inicio()
    {
    inteiro a1, r, n, i, termo, soma

    
    escreva("Digite o primeiro termo, a razão e a quantidade de termos: ")
    leia(a1, r, n)

    soma = 0
        
    para(i = 0; i < n; i++){  
    termo = a1 + (i * r)
    soma = soma + termo}
        
    escreva(soma, "\n")
    }
}
