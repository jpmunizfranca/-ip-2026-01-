package main
import "fmt"

func main() {
  var a, b, media, pesoa, pesob float64
  fmt.Println("Me fale a primeira nota.")
  fmt.Scan(&a)
  fmt.Println("Agora me fale o peso dela.")
  fmt.Scan(&pesoa)
  fmt.Println("Me fale a segunda nota.")
  fmt.Scan(&b)
  fmt.Println("Agora me fale o peso dela.")
  fmt.Scan(&pesob)
 
   
  media = (a * pesoa + b * pesob) / (pesoa + pesob)
  fmt.Println("Sua média é:", media)
  
}