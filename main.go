package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	var contaDoGuilherme contas.ContaCorrente = contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Guilherme", CPF: "123.478.258-52", Profissao: "Gerente"},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	contaDoGuilherme.Depositar(125.50) 

	PagarBoleto(&contaDoGuilherme, 60)

	contaDaMaria := contas.ContaPoupanca{}

	fmt.Println(contaDaMaria)

	clienteJulia := clientes.Titular{Nome: "Julia", CPF: "457.180.448-21", Profissao: "Especialista fullStack and Data"}

	contaDaJulia := contas.ContaCorrente{
		Titular:       clienteJulia,
		NumeroAgencia: 589,
		NumeroConta:   123486,
	}

	contaDaJulia.Depositar(100000.00)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = clientes.Titular{Nome: "Cris", CPF: "145.145.158-47", Profissao: "PO"}
	contaDaCris.Depositar(500)

	fmt.Println(contaDaCris.Sacar(200.))
	fmt.Println(contaDaJulia.Depositar(5000.00))

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaJulia)
	fmt.Println(*contaDaCris)

	status := contaDaCris.Transferir(100, &contaDoGuilherme)
	fmt.Println(status, contaDaCris, contaDoGuilherme)
}
