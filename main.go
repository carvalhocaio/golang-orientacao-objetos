package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (cc *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= cc.saldo

	if podeSacar {
		cc.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (cc *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		cc.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", cc.saldo
	} else {
		return "O valor do depósito é inválido!", cc.saldo
	}
}

func main() {
	contaMartin := ContaCorrente{}
	contaMartin.titular = "Martin"
	contaMartin.saldo = 500

	fmt.Println(contaMartin.saldo)
	fmt.Println(contaMartin.Depositar(200))
	fmt.Println(contaMartin.saldo)
}
