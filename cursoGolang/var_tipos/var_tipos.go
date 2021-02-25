package main

import (
	"fmt"
	"reflect"
)

func main(){
	var nome string = "Carol"
	var idade = 27 //inferência do tipo de variável int, string, bool
	versao := 1.1 //declaracao e inferência de variável ao mesmo tempo e de forma direta

	fmt.Println("Olá", nome)
	fmt.Println("Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao)) //o pacote reflect.TypeOf() permite descobrir o tipo de determinada variável
}