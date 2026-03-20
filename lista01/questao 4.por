programa {
  funcao inicio() {
    real salario
    real kw
    real custo
    real valor
    real desconto 

escreva("Me infome o valor do salário mínimo atual e a quantidade consumida de kW, respectivamente.")
leia(salario) leia(kw)

custo = salario * 7 / 1000
valor = custo * kw
desconto = valor * 9 / 10

escreva("ENTRADA")
escreva("\n") escreva(salario)
escreva("\n") escreva(kw)

escreva("\nSAÍDA")
escreva("\n") escreva("custo por kW: ") escreva("R$") escreva(custo)

escreva("\nCusto do consumo: ") escreva("R$") escreva(valor)

escreva("\nCusto com desconto: ") escreva("R$") escreva(desconto)

  }
}
