package contas

import "github.com/carvalhocaio/golang-orientacao-objetos/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
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

func (cc *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < cc.saldo && valorDaTransferencia > 0 {
		cc.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (cc *ContaCorrente) ObterSaldo() float64 {
	return cc.saldo
}
