package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":           1200.90,
		"GAbriel":        1400.20,
		"Arthur Migurel": 980.54,
	}

	funcsESalarios["Priscilla"] = 1340.90
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}