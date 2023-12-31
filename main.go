package main

import (
	"fmt"

	"github.com/carvalhocaio/golang-orientacao-objetos/clientes"
	"github.com/carvalhocaio/golang-orientacao-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

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
	contaMartin.Depositar(100)
	PagarBoleto(&contaMartin, 60)
	fmt.Println(contaMartin.ObterSaldo())
}
