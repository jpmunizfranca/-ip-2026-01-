package main
import "fmt"

func main() {
  var x int 
  fmt.Println("Me informe um valor inteiro maior que zero.")
  fmt.Scan(&x)
  fmt.Println("O fatorial desse número é", fatorial(x))
}
func fatorial(x int) int {
    if x == 0 {
        return 1
    }
    return x * fatorial(x-1)
}