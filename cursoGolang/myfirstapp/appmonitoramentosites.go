package main

import (
	"fmt"
	"net/http"
	"os"
)

func exibirIntroducao() {
	fmt.Println("Insira seu nome: ")
	nome := lerNome()
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println("Escolha uma opção para dar início ao programa: ")
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
}

func iniciarMonitoramento() {
	fmt.Println("Insira o site que deseja fazer o monitoramento: ")
	//array
	/*
		var site [3]string
		site[0] = lerSite()
		site[1] = lerSite()
		site[1] = lerSite()

		resp, _ := http.Get(site[i]) //o pacote net/http permite fazer requisições de acesso a web (Get, Post...) e além da resposta, ela tbm retorna um possível erro, que podemos ignorar
		fmt.Println(resp)

		if resp.StatusCode == 200 {
				fmt.Println("Site", site[i], "foi carregado com sucesso!")
			} else {
				fmt.Println("Site", site[i], "com probelmas, Status Code: ", resp.StatusCode)
			}
	*/

	//slice
	site := []string{lerSite(), lerSite(), lerSite()}

	//for range - passado a posicao e o que ela contém
	for i, sites := range site {
		fmt.Println("Testando site: ", i, ": ", sites)
	
		resp, _ := http.Get(site[i])

		if resp.StatusCode == 200 {
			fmt.Println("Site", site[i], "foi carregado com sucesso!")
		} else {
			fmt.Println("Site", site[i], "com probelmas, Status Code: ", resp.StatusCode)
		}
	}

	//for comum- varrendo o slice até o tamanho dele, passando somente o conteúdo dele
	/*for i := 0; i < len(site); i++ {
		//fmt.Println(site[i])
		resp, _ := http.Get(site[i])

		if resp.StatusCode == 200 {
			fmt.Println("Site", site[i], "foi carregado com sucesso!")
		} else {
			fmt.Println("Site", site[i], "com probelmas, Status Code: ", resp.StatusCode)
		}
	}*/
}

func main() {
	exibirIntroducao()

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
