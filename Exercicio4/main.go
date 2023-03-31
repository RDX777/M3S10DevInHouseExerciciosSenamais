package main

import (
	"fmt"
	"math"
)

type Carteira interface {
	enviar(bitcoin float64, endereco string)
	receber(bitcoin float64)
	consultarSaldo() float64
}

type Endereco struct {
	chavePublica string
	chavePrivada string
	saldo        float64
}

func (e *Endereco) enviar(bitcoin float64, endereco string) {
	if e.saldo < bitcoin {
		fmt.Println("Saldo insuficiente")
		return
	}
	e.saldo -= bitcoin
	fmt.Printf("Enviando %.8f BTC para o endereço %s\n", bitcoin, endereco)
}

func (e *Endereco) receber(bitcoin float64) {
	e.saldo += bitcoin
	fmt.Printf("Recebendo %.8f BTC\n", bitcoin)
}

func (e *Endereco) consultarSaldo() float64 {
	return math.Round(e.saldo*100000000) / 100000000
}

func enviarBitcoin(c Carteira, bitcoin float64, endereco string) {
	c.enviar(bitcoin, endereco)
}

func main() {

	endereco := Endereco{
		chavePublica: "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
		chavePrivada: "5Jh22WnveEaL6mK1emvC8LLTgT6UZj6UZDFNnxEfwXNBrN6LUTn",
		saldo:        1.23456789,
	}

	enviarBitcoin(&endereco, 0.12345678, "3KJmDCM1NiTujxNBhZP6iZf6Uc9XHJ9XUy")

	fmt.Printf("Saldo atual: %.8f BTC\n", endereco.consultarSaldo())
}
