package main

import f "fmt"

func main() {
	var idades [50]int
	f.Println("Digite 50 idades:")
	for i := range idades {
		f.Scan(&idades[i])
	}

	freq := map[int]int{}
	for _, v := range idades {
		freq[v]++
	}

	moda, maxFreq := 0, 0
	for idade, cnt := range freq {
		if cnt > maxFreq {
			moda, maxFreq = idade, cnt
		}
	}
	f.Printf("Moda: %d (apareceu %d vezes)\n", moda, maxFreq)
}
