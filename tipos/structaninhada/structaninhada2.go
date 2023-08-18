package main

type item struct {
	produto1ID int
	qtde1 int
	preco1 float64
}

type pedido1 struct {
	user1ID int
	itens1 []item
}

func (p pedido1) valorTotal1() float64 {
	total1 := 0.0
	for _, item := range p.itens1 {
		total1 += item.preco1 * float64(item.qtde1)
	}
	return total1
}

func main() {
	
}