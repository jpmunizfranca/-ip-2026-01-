package main
import "fmt"

func main() {
    var x int 
    fmt.Println("Me vale o valor de x.")
    fmt.Scan(&x)
    if x > 20 && x < 90 {
        fmt.Println("Esse número está entre 20 e 90.")
    } else {
        fmt.Println("Esse numero não está entre 20 e 90.")
    }
    
}