package main

import "fmt"

func main() {
	// var aprovados map[int]
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678912] = "Maria"
	aprovados[43230948457] = "João"
	aprovados[99845887453] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678912])
	delete(aprovados, 12345678912)
	fmt.Println(aprovados[12345678912])
}