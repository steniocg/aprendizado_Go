package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15214.90,
			"Gabriel Mongo":  12900.91,
		},
		"B": {
			"Bruno Arthur": 8090.76,
			"Babi buxa":    9012.09,
		},
		"P": {
			"Priscilla Guedes": 1209.04,
			"Papai José":       12390.09,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}