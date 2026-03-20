programa {questao 1
  funcao inicio() {
    real valor1, valor2, valor3, media

    escreva("Quais são as suas notas? ")
    leia(valor1, valor2, valor3)

    media =(valor1 + valor2 + valor3) / 3 

escreva("ENTRADA")
escreva("\n")
    escreva(valor1,  -  valor2,   -  valor3)

escreva("\n")
escreva("SAÍDA")
    escreva("\nSua média é: ") 
    escreva(media)

    se(media < 6){
      escreva ("\nREPROVADO")
    }

    se(media >= 6){
      escreva("\nAPROVADO")

    }
    
  }
}
