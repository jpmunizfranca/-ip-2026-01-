package main
import "fmt"

func main() {
    var x, y, z int
    fmt.Println("Me fale valores inteiros diferentes para x, y, z.")
    fmt.Scan(&x, &y, &z)
    fmt.Println("O maior valor entre eles é", maior(x, y, z))
}
func maior(x, y, z int) int {
    if x > y && x > z {
        return x
    }
    if y > x && y > z {
        return y
    } else {
        return z
    }