package main

import "fmt"

type Veiculo struct {
	marca  string
	modelo string
}

type Carro struct {
	Veiculo
	numeroDePortas int
}

type Moto struct {
	Veiculo
	cilindradas int
}

type Revisao interface {
	fazerRevisao()
}

func (c Carro) fazerRevisao() {
	fmt.Printf("Fazendo revisão do carro %s %s com %d portas...\n", c.marca, c.modelo, c.numeroDePortas)
}

func (m Moto) fazerRevisao() {
	fmt.Printf("Fazendo revisão da moto %s %s com %d cilindradas...\n", m.marca, m.modelo, m.cilindradas)
}

func agendarRevisao(r Revisao) {
	r.fazerRevisao()
}

func main() {
	carro := Carro{
		Veiculo:        Veiculo{"Volkswagen", "Golf"},
		numeroDePortas: 4,
	}

	moto := Moto{
		Veiculo:     Veiculo{"Honda", "CB500"},
		cilindradas: 500,
	}

	agendarRevisao(carro)
	agendarRevisao(moto)
}
