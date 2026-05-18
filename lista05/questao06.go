package main

import f "fmt"

func main() {
	var v [100]int
	for i := range v {
		v[i] = 100 - i
	}
	f.Println(v)
}
