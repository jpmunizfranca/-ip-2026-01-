programa
{
    inclua biblioteca Matematica --> mat
    inclua biblioteca Tipos --> tip

    funcao inicio()
    {
    inteiro n, k
    real S

    escreva("Digite um número inteiro positivo maior que 1: ")
    leia(n)

      
    se(n <= 1){
    escreva("Numero invalido!\n")}
  
    senao{ 
     S = 0.0
            
    para(k = 1; k <= n; k++){
            
     S = S + (1.0 / tip.inteiro_para_real(k))}
            
    S = mat.arredondar(S, 6)
     
    escreva(S, "\n")
        }
    }
}
