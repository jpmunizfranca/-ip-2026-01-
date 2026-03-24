package main
import "fmt"

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("Me fale o valor do lado A")
	fmt.Scan(&a) 
	fmt.Println("Me fale o valor do lado B")
	fmt.Scan(&b)
	fmt.Println("Me fale o valor do lado C")
	fmt.Scan(&c)
	
	if a > b + c || b > c + a || c > b + a {
	fmt.Println("Esse triângulo não existe.")
	} else if a == b && b == c {
	fmt.Println("Esse triângulo é equilátero.")
	} else if a == b || b == c || a == c {
	fmt.Println("Esse triângulo é isósceles.")
	} else { 
	fmt.Println("Esse trinângulo é escaleno.")
	}
	
	
}