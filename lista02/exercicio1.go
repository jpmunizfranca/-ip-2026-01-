package main
import "fmt"

func main() {
  var a int
  fmt.Println("Me fale um valor inteiro maior ou igual a 2.")
  fmt.Scan(&a)
  if a % 2 == 0 {
     fmt.Println("Esse número é divisível por 2.")
  }
  if a % 3 == 0 {
      fmt.Println("Esse número é divisível por 3.")
  }
}