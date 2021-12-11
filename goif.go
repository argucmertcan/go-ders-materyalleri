package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çift sayı")
		} else if i%3 == 0 {
			fmt.Println(i, "3 e tam bölündü")
		} else {
			fmt.Println(i, "tek sayı")
		}
	}
}
