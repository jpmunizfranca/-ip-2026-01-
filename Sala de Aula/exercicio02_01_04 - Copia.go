package main
import "fmt"

func main() {
  var x, y, z float64
  fmt.Println("Me fale três valores reais.")
  fmt.Scan(&x, &y, &z)
  fmt.Printf("A média entre eles é %.2f", media(x, y, z))
}
func media(x, y, z float64) float64 {
    media := (x + y + z) / 3
    return media
}