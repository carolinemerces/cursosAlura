package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func exibirIntroducao() {
	fmt.Println("Insira seu nome:")
	nome := lerNome()
	versao := 1.2
	fmt.Println("")
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão:", versao)
	fmt.Println("Escolha uma opção para dar início ao programa:")
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

func lerSitesdoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") //o pacote os, com a função Open, permite abrir um arquivo, porém mostrando seu endereço
	//arquivo, err := ioutil.ReadFile("sites.txt") //o pacote ioutil, com a função ReadFile, permite abrir e ler um arquivo, retornando um array de bytes, porém desta forma é mais fácil convertar uma uma string
	//fmt.Println(string(arquivo)) //convertendo esse arquivo que está retornando um array de bytes em string

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo) //o pacote bufio, com a função NewReader, permite abrir e ler um arquivo e através de outra função lê-lo e convertê-lo para string

	//utilização do for para repetir a mesma ação até que ele retorne um erro de que não há mais linhas para ler (End On File - EFO)
	for {
		linha, err := leitor.ReadString('\n') //a função ReadSring, converte os bytes em string, percorrendo cada byte até onde preciso que ele leia. No caso até a quebra de linha, representada pelo \n, com '', pois se trata de byte
		linha = strings.TrimSpace(linha)      //a função TrimSpace, do pacote de strings, permite retirar os espaços e pular uma linha a cada execução da função acima
		//fmt.Println(linha)

		sites = append(sites, linha) //ao invés de imprimir a linha, colocamos ela dentro do nosso slice de sites

		if err == io.EOF { //quando o erro for igual a não há mais linhas para leitura, quero que ele pare a execução do for
			//fmt.Println("Ocorreu um erro: ", err)
			break
		}
	}

	//fmt.Println(sites)

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
			fmt.Println("Testando site:", i, ":", sites)
			//fmt.Println("")

			resp, err := http.Get(site[i])

			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Site", site[i], "foi carregado com sucesso!")
				fmt.Println("")
			} else {
				fmt.Println("Site", site[i], "com problemas, Status Code:", resp.StatusCode)
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
