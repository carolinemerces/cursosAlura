package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 1
const delay = 5
const versao = 1.3

func exibirIntroducao() {
	fmt.Println("Insira seu nome:")
	nome := lerNome()
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

func registraLog(site string, status bool) { //função para abrir e criar um arquivo para registrar os logs dos sites (data, horário e status)
	//arquivo, err := os.Open("log.txt")//irá abrir o arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //a função openFile, permite além de abrir um arquivo, criar caso ele não exista e editar, escrevendo (através da flag, que é um dos parâmetros exigidos para essa função) e adicionar um número para permissão de acesso ao arquivo
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLog() { //essa função vai abrir, ler e imprimir o arquivo, sem a necessidade de fechá-lo, por conta do pacote ioutil
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
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
	arquivo.Close() //são boa práticas sempre que abrir um arquivo, fechá-lo para ele possa ser usar em outras situações dentro do programa
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
				registraLog(site[i], true)
				fmt.Println("")
			} else {
				fmt.Println("Site", site[i], "com problemas, Status Code:", resp.StatusCode)
				registraLog(site[i], false)
				fmt.Println("")
			}
		}
		time.Sleep(delay * time.Second) //o pacote time, possui a função sleep que permite colocar um tempo para que a operação seja executada
	}
}

func main() {
	exibirIntroducao()
	//lerSitesdoArquivo()
	//registraLog("site-false", false)

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
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //o pacote os, contém a função exit, que permite sair do programa, de acordo com o parâmetro 0 passado a ele
		default:
			fmt.Println("Comando não reconhecido!")
			os.Exit(-1) //assim como também permite informar que algo deu errado no programa, de acordo com o parâmetro -1 passado a ele
		}
	}
}
