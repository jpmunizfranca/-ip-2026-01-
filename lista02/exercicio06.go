package main
import "fmt"

func main() {
    var a int
    var b int
    fmt.Println("Me informe um número inteiro a.")
    fmt.Scan(&a)
    fmt.Println("Me informe um número inteiro b.")
    fmt.Scan(&b)
    if a % b == 0 {
        fmt.Println("O número a é divisível por b.")
    } else {
        fmt.Println("O número a não é divisível por b.")
    }
}