package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaMartin := ContaCorrente{
		titular:       "Martin",
		numeroAgencia: 555,
		numeroConta:   112233,
		saldo:         100.20,
	}

	fmt.Println(contaMartin)

	contaHailey := ContaCorrente{"Hailey", 333, 111333, 200}
	fmt.Println(contaHailey)
}
