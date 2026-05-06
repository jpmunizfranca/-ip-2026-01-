package main

import f "fmt"

func main() {
	var v [100]int
	impar := 1
	for i := range v {
		v[i] = impar
		impar += 2
	}
	f.Println(v)
}
