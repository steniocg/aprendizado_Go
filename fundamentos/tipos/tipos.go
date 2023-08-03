package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros

	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positios)... uint8 uint16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64

	i1 := math.MaxInt64
	fmt.Println("o valor maximo int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento de tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo da váriavel bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá Stênio!!!"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	stenio`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	char := 'a'
	fmt.Println("o tipo de char é", reflect.TypeOf(char))
	

}