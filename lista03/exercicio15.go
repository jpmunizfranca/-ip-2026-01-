package main
import f "fmt"

func main() {
	var n int
	f.Println("Digite o valor de n-ésimo termo:")
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		f.Println(i * i)
	}
