programa {
  inclua biblioteca Matematica
  funcao inicio() {
    inteiro x 
    inteiro y
    real r, c
    escreva("Me informe o número de horas locadas.")
leia(x)
y = x / 3
r = x % 3
c = (10 * y) + (5 * r)
c = Matematica.arredondar(c, 2)
escreva("O valor a pagar é = ") escreva(c)
escreva("\n")

  }
}
