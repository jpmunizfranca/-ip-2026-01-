programa {
  funcao inicio() {
    real a, b, c, d
    real m
    escreva("Me informe os valores de a, b, c e d da matriz quadrada bidimensional.")
    leia(a, b, c, d)
    m = (a * d) - (c * b)
    escreva("O valor do determinante da matriz é: ") escreva(m) escreva("\n")
  }
}
