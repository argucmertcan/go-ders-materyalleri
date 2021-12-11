package main

import "fmt"

func main() {
	x := []float64{15, 23, 435, 56, 33}
	fmt.Println(ortalama(x))
}

func ortalama(x []float64) float64 {
	var toplam float64 = 0
	for i := 0; i < 5; i++ {
		toplam += x[i]
	}

	return toplam / float64(len(x))
}
