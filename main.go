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

func (cc *ContaCorrente) Tranferir(
	valorDaTransferencia float64,
	contaDestino *ContaCorrente,
) bool {
	if valorDaTransferencia < cc.saldo && valorDaTransferencia > 0 {
		cc.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaMartin := ContaCorrente{titular: "Martin", saldo: 300}
	contaHailey := ContaCorrente{titular: "Hailey", saldo: 100}

	status := contaMartin.Tranferir(150, &contaHailey)

	fmt.Println(status)

	fmt.Println(contaMartin)
	fmt.Println(contaHailey)
}
