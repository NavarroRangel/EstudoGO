package main

import "fmt"

type ContaCorrente struct {
	titular 	string
	numeroAgencia	 int
	numeroConta 	int
	saldo 	float64
}

func (c *ContaCorrente)Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 &&valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaNavarro := ContaCorrente {}
	contaNavarro.titular = "naarro"
	contaNavarro.saldo =500
	fmt.Println(contaNavarro)
	
	
	fmt.Println(contaNavarro.Sacar(200))
	fmt.Println(contaNavarro.saldo)
	
}
