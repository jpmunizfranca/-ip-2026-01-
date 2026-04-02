package main
import ("fmt"; "math")
func main() {
  var x, y float64
  fmt.Println("Me fale um número x.")
  fmt.Scan(&x)
  if x > 0 {
     fmt.Printf("O valor de sua raiz quadrada é : %.2f", math.Sqrt(x))
  } else {
      y = x*x
      fmt.Printf("Seu quadrado é: %.2f", y)
  }
}