package main

import "fmt"

func media(numero ...float64) float64 {
	total := 0.0
	for _, num := range numero {
		total += num
	}
	return total / float64(len(numero))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 6.5, 9.0))
}