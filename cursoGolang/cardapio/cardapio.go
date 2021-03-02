package main

import "fmt"

func main() {
	var prato string
	fmt.Println("Digite seu prato preferido...")
	fmt.Println("P - Pizza")
	fmt.Println("H - Hambúrguer")
	fmt.Println("B - Bife com fritas")
	fmt.Println("S - Salada Caesar")
	fmt.Println("F - Salada de Frutas")
	fmt.Println("E - Estrogonofe")
	fmt.Println("O - Outros")
	fmt.Scan(&prato)

	switch prato {
	case "B":
		fmt.Println("Com batatas Palito ou Noisete?")
	case "H":
		fmt.Println("Com Queijo ou com Ovo?")
	case "P":
		fmt.Println("Calabresa ou Napolitana?")
	case "S":
		fmt.Println("Alface ou Rúcula?")
	case "F":
		fmt.Println("Kiwi ou Frango?")
	case "E":
		fmt.Println("Carne ou Frango?")
	case "O":
		fmt.Println("Não gostou de nosso cardápio?")
	default:
		fmt.Println("Não entendi seu paladar...")
	}
}
