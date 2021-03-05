package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func exibirIntroducao() {
	fmt.Println("Insira seu nome: ")
	nome := lerNome()
	versao := 1.2
	fmt.Println("")
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println("Escolha uma opção para dar início ao programa: ")
	fmt.Println("")
}

func lerNome() string {
	var nome string
	fmt.Scan(&nome)
	return nome
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func lerSite() string {
	var site string
	fmt.Scan(&site)
	return site
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
	fmt.Println("")
}

func lerSitesdoArquivo() []string{
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(arquivo)
	return sites
}

func iniciarMonitoramento() {
	//fmt.Println("Insira o site que deseja fazer o monitoramento: ")
	fmt.Println("")

	//sites em um arquivo
	site := lerSitesdoArquivo()

	//slice
	//site := []string{lerSite(), lerSite()}

	//for range - passado a posicao e o que ela contém
	for i := 0; i < monitoramento; i++ { //colocar um tempo para repetir a operação
		for i, sites := range site {
			fmt.Println("Testando site: ", i, ": ", sites)
			fmt.Println("")

			resp, err := http.Get(site[i])

			if err != nil {
				fmt.Println("Ocorreu um erro: ", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Site", site[i], "foi carregado com sucesso!")
				fmt.Println("")
				} else {
				fmt.Println("Site", site[i], "com probelmas, Status Code: ", resp.StatusCode)
				fmt.Println("")
			}
		}
		time.Sleep(delay * time.Second) //o pacote time, possui a função sleep que permite colocar um tempo para que a operação seja executada
	}
}

func main() {
	exibirIntroducao()
	lerSitesdoArquivo()

	for {
		exibirMenu()

		comando := lerComando()
		fmt.Printf("O comando escolhido foi %d\n", comando)

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //o pacote os, contém a função exit, que permite sair do programa, de acordo com o parâmetro 0 passado a ele
		default:
			fmt.Println("Comando não reconhecido!")
			os.Exit(-1) //assim como também permite informar que algo deu errado no programa, de acordo com o parâmetro -1 passado a ele
		}
	}
}
