package main

import (
	"fmt"

	"github.com/carvalhocaio/golang-orientacao-objetos/clientes"
	"github.com/carvalhocaio/golang-orientacao-objetos/contas"
)

func main() {
	clienteMartin := clientes.Titular{
		Nome:      "Martin",
		CPF:       "01234567890",
		Profissao: "Desenvolvedor",
	}

	contaMartin := contas.ContaCorrente{
		Titular:       clienteMartin,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	fmt.Println(contaMartin)
	fmt.Println(contaMartin.Depositar(120))
	fmt.Println(contaMartin.ObterSaldo())

	clienteHailey := clientes.Titular{
		Nome:      "Hailey",
		CPF:       "01234567890",
		Profissao: "Designer",
	}

	contaHailey := contas.ContaPoupanca{
		Titular:       clienteHailey,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	fmt.Println(contaHailey)
	fmt.Println(contaHailey.Depositar(200))
	fmt.Println(contaHailey.ObterSaldo())
}
