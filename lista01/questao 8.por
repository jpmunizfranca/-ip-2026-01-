programa {
  funcao inicio() {
    real r, h 
    real at, ac, al, custo
    escreva("Me fale o raio da lata.")
    leia(r)
    escreva("\nMe fale a altura da lata.")
    leia(h)
    al = 3.14159 * 2 * r * h
    ac = 3.14159 * r * r
    at = 2 * ac + al
    custo = 100 * at
escreva("O valor do custo é = ") escreva(custo)
    

  }
}
