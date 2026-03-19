programa {
  funcao inicio() {
    
real r1, r2, r3
inteiro conta
real consumo  
cadeia consumidor 
cadeia r, c, i 

escreva("Me informe o número da sua conta, seu consumo de água e o seu tipo de consumidor, residencial (r), comercial (c) ou industrial (i).")
leia(conta) leia(consumo) leia(consumidor)

  r1 = ( 0.05* consumo ) + 5
  r2 = ( consumo - 80 ) * (0.25) + 500
  r3 = ( consumo - 100) * ( 0.04) + 800

  escreva("CONTA = ") escreva(conta)


se(consumidor == "r"){
  escreva("\nVALOR DA CONTA = ") escreva(r1)}

  se(consumidor == "c"){
    escreva("\nVALOR DA CONTA = ") escreva(r2)}

    se(consumidor == "i"){
      escreva("\nVALOR DA CONTA = ") escreva(r3)}
    }
  }
}
}

















  }
}
