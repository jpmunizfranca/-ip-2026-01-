package main
import(
	"fmt"
)
func main () {

var x int
fmt.Println("Escreva um valor inteiro.")
fmt.Scan(&x)
if x > 0 {
    fmt.Println("Esse valor é positivo.")
} else if x < 0 {
    fmt.Println("Esse valor é negativo.")
} else {
    fmt.Println("Esse valor é nulo.")
}

}