package main

import "fmt"

func main(){
	nome := "Nome"
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão: ", versao)

	//Menu
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	//Pacote scan permite capturar o que o usuário irá escrever
	var comando int
	//fmt.Scanf("%d", &comando) //a funcao scan espera receber um endereco da variável e não a variável em si, por isso deve-se usar o &, que indica o endereço da variável
	 
	fmt.Scan(&comando) //outra forma "mais limpa" e por inferência
	
	/*
	if comando == 1{
		fmt.Println("Monitorando...")
	} else if comando == 2{
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conhece esse comando!")
	}*/

	switch comando {
	case 1: 
	fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conhece esse comando!")
	}

	
}
