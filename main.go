package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var contaMartin *ContaCorrente
	contaMartin = new(ContaCorrente)

	contaMartin.titular = "Martin"
	contaMartin.numeroAgencia = 1
	contaMartin.numeroConta = 123
	contaMartin.saldo = 100

	fmt.Println(*contaMartin)
}
