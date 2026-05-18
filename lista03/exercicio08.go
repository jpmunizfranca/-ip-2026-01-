package main

import f "fmt"

func main() {
	soma1 := 0
	for i := 1; i <= 20; i++ {
		f.Println(i)
		soma1 += i
	}
	f.Println("A soma é :", soma1)
}
