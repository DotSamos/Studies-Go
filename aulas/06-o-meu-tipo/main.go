package main

import "fmt"

/*
	É possível criar tipos customizados no go utilizando a keyword "type"
	type <nome> <tipo-base>
*/

// criamos o tipo "customString" utilizando o tipo padrão string como base
type customString string

// criamos uma variável onde o tipo dela é o "customString", e o valor dela não deixa de ser uma string padrão devido a sua base
var textOne customString = "Um texto aqui"

// criamos uma outra variável aqui só que com o tipo string, que é a base do tipo "customString"
var textTwo string = "Um texto aqui"

/*
	apesar do valor das variáveis serem iguais elas não são totalmente iguais, afinal o seu tipo
	é diferente
*/

func main() {
	fmt.Printf("Texto com um tipo custom: %s (%T)\n", textOne, textOne)

	// textTwo = textOne /* Erro, pois o tipo textTwo é uma string, em quanto textOne é um "customString" */

	// para isso funcionar é necessário fazer uma conversão
	textTwo = string(textOne)

	fmt.Printf("Texto convertido: %s (%T)\n", textTwo, textTwo)
}
