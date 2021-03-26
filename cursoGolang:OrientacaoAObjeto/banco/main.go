package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) (status string) {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Slado insuficiente!"
}

func (c *ContaCorrente) Depositar(deposito float64) (status string, saldo float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Depósito realizado com sucesso!", c.saldo
	}
	return "O valor do depósito menor que zero!", c.saldo
}

func (c *ContaCorrente) Transferir(transferencia float64, contaDestino *ContaCorrente) (status bool) {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	}
	return false
}

func main() {
	conta1 := ContaCorrente{"Caroline", 589, 123456, 125.5}
	fmt.Println(conta1)

	conta2 := new(ContaCorrente)
	conta2.titular = "Camila"
	conta2.saldo = 230.99
	fmt.Println(*conta2)

	// saque := 20.00
	// conta2.saldo = conta2.saldo - saque
	// fmt.Println(conta2.saldo)

	fmt.Printf("Seu saldo: %2.f\n", conta2.saldo)
	fmt.Println(conta2.Sacar(50.00))

	fmt.Printf("Seu saldo: %2.f\n", conta2.saldo)
	fmt.Println(conta2.Depositar(60.00))

	status := conta1.Transferir(45.87, conta2)
	fmt.Println(status)
	fmt.Println(conta1.saldo, conta2.saldo)
}
