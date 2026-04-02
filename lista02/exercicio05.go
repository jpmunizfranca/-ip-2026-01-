package main
import "fmt"
func main() {
   var x int
   fmt.Println("Me informe um número inteiro.")
   fmt.Scan(&x)
   if x % 5 == 0 {
       fmt.Println("Esse número é divisível por 5.")
   } else {
       fmt.Println("Esse número não é divisível por5.")
   }
}