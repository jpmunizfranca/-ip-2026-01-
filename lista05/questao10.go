package main

import f "fmt"

func main() {
	var fib [50]int
	fib[0], fib[1] = 1, 1
	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	f.Println(fib)
}
