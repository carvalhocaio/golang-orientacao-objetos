package main

import (
	"fmt"

	"github.com/carvalhocaio/golang-orientacao-objetos/contas"
)

func main() {
	contaMartin := contas.ContaCorrente{Titular: "Martin", Saldo: 300}
	contaHailey := contas.ContaCorrente{Titular: "Hailey", Saldo: 100}

	status := contaMartin.Tranferir(150, &contaHailey)

	fmt.Println(status)

	fmt.Println(contaMartin)
	fmt.Println(contaHailey)
}
