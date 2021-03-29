package main

import (
	"fmt"
)

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

	fmt.Printf("Seu saldo: %2.f\n", conta2.Saldo)
	fmt.Println(conta2.Sacar(50.00))

	fmt.Printf("Seu saldo: %2.f\n", conta2.Saldo)
	fmt.Println(conta2.Depositar(60.00))

	status := conta1.Transferir(45.87, conta2)
	fmt.Println(status)
	fmt.Println(conta1.Saldo, conta2.Saldo)
}
