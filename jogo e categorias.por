programa {questao 2
  funcao inicio() {
    real valor1, valor2, valor3, valor4
    real renda
real p1
real p2
real p3
real p4
real publico

    

escreva(" Me informe o número de pessoas em cada categoria, sendo elas Popular, Geral, Arquibancada e Cadeira, respectivamente. ")
leia(valor1, valor2, valor3, valor4)

renda = (valor1) + (valor2 * 5) + (valor3 * 10 ) + (valor4 * 20)
publico = (valor1) + (valor2) + (valor3) + (valor4)

p1 = valor1 / publico
p2 = valor2 / publico
p3 = valor3 / publico
p4 = valor4 / publico

escreva("\nENTRADA ")
escreva("\nO público total do jogo é: ")
escreva(publico)
escreva("\nPorcentegem de pessoas na categoria Popular: ")
escreva(p1)

escreva("\nPorcentagem de pessoas na categoria Geral: ")
escreva(p2)

escreva("\nPorcentagem de pessoas na categoria Arquibancada: ")
escreva(p3)

escreva("\nPorcentagem de pessoas na categoria Cadeiras: ")
escreva(p4)

escreva("\nSAÍDA")
escreva("\nA renda total do jogo é: ")
escreva(renda)

 


    
  }
}
