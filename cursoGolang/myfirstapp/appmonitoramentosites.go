package main

import "fmt"

func exibirIntroducao() {
	fmt.Println("Insira seu nome: ")
	nome := lerNome()
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println("Escolha uma opção para dar início ao programa: ")
}

func lerNome() string{
var nome string
fmt.Scan(&nome)
return nome
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibirMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func sair(){
	
}

func main() {
	exibirIntroducao()
	exibirMenu()

	comando := lerComando()
	fmt.Printf("O comando escolhido foi %d", comando)

}
