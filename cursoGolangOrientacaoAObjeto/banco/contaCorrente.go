package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) (status string) {
	podeSacar := saque > 0 && saque <= c.Saldo
	if podeSacar {
		c.Saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Slado insuficiente!"
}

func (c *ContaCorrente) Depositar(deposito float64) (status string, Saldo float64) {
	if deposito > 0 {
		c.Saldo += deposito
		return "Depósito realizado com sucesso!", c.Saldo
	}
	return "O valor do depósito menor que zero!", c.Saldo
}

func (c *ContaCorrente) Transferir(transferencia float64, contaDestino *ContaCorrente) (status bool) {
	if transferencia < c.Saldo && transferencia > 0 {
		c.Saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	}
	return false
}
