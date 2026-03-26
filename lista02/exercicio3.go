package main
import(
    "fmt"
    )
func main () {
var x, y int
var soma1, soma2, soma3 int
fmt.Println("Digite o valor de x.")
fmt.Scan(&x)
fmt.Println("Digite o valor de y.")
fmt.Scan(&y)
soma1 = x + y
if soma1 > 20 {
    soma2 = soma1 + 8
    fmt.Println("Valor a ser apresentado: ", soma2)
} else{
    soma3 = soma1 - 5
    fmt.Println("Valor a ser apresentado:", soma3)
}
}
