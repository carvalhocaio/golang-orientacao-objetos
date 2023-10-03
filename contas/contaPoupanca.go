package contas

import "github.com/carvalhocaio/golang-orientacao-objetos/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (cc *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= cc.saldo

	if podeSacar {
		cc.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (cc *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		cc.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", cc.saldo
	} else {
		return "O valor do depósito é inválido!", cc.saldo
	}
}

func (cc *ContaPoupanca) ObterSaldo() float64 {
	return cc.saldo
}
