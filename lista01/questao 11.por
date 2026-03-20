programa {
  funcao inicio() {
    real x
    escreva("Me informe um valor.")
    leia(x)
    se(x%3 == 0 e x%5 == 0){
      escreva("Esse número é divsível.")}
      senao{
        escreva("Esse número não é divisível.")
      }
  }
}
