package main

import (
	f "fmt"
	"math"
)

func main() {
	var x float64
	f.Print("Digite o valor de x (em radianos): ")
	f.Scan(&x)

	coseno := 0.0
	sinal := 1.0
	fatorial := 1.0
	potencia := 1.0

	for i := 0; i < 20; i++ {
		if i == 0 {
			coseno += 1.0
		} else {
			fatorial *= float64(2*i-1) * float64(2*i)
			potencia *= x * x
			sinal *= -1
			coseno += sinal * potencia / fatorial
		}
	}

	diferenca := coseno - math.Cos(x)
	f.Printf("a) Cosseno calculado pela série: %.6f\n", coseno)
	f.Printf("b) Diferença em relação a math.Cos(x): %.6f\n", diferenca)
}
