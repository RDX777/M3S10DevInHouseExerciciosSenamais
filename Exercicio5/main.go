package main

import (
	"fmt"
	"math/rand"
)

type Carteira interface {
	enviar(valor float64, destino string) error
	receber(valor float64) error
	saldo() float64
}

type Endereco struct {
	chavePublica    string
	chavePrivada    string
	saldoDisponivel float64
}

func (e *Endereco) enviar(valor float64, destino string) error {
	if e.saldoDisponivel < valor {
		return fmt.Errorf("Saldo insuficiente para realizar a transação")
	}
	e.saldoDisponivel -= valor
	fmt.Printf("%v Bitcoins foram enviados para o endereço %v\n", valor, destino)
	return nil
}

func (e *Endereco) receber(valor float64) error {
	e.saldoDisponivel += valor
	fmt.Printf("%v Bitcoins foram recebidos na carteira\n", valor)
	return nil
}

func (e *Endereco) saldo() float64 {
	return e.saldoDisponivel
}

type Minerador interface {
	minerar() float64
}

type MineradorDeBitcoin struct {
	dificuldade int
}

func (m *MineradorDeBitcoin) minerar() float64 {
	valor := rand.Float64() * float64(m.dificuldade)
	fmt.Printf("%v Bitcoins foram minerados\n", valor)
	return valor
}

func enviarBitcoin(c Carteira, valor float64, destino string) error {
	err := c.enviar(valor, destino)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	endereco := Endereco{
		chavePublica:    "abc123",
		chavePrivada:    "def456",
		saldoDisponivel: 10,
	}

	minerador := MineradorDeBitcoin{
		dificuldade: 10,
	}

	valorMinerado := minerador.minerar()
	endereco.receber(valorMinerado)

	err := enviarBitcoin(&endereco, 5, "xyz789")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Saldo disponível: %v Bitcoins\n", endereco.saldo())
}
