package contas

import "github.com/carvalhocaio/golang-orientacao-objetos/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (cc *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= cc.Saldo

	if podeSacar {
		cc.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (cc *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		cc.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", cc.Saldo
	} else {
		return "O valor do depósito é inválido!", cc.Saldo
	}
}

func (cc *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < cc.Saldo && valorDaTransferencia > 0 {
		cc.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
