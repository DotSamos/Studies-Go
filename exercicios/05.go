package main

import "fmt"

/*

	Utilizando a solução do exercício anterior:

	1. EM package-level scope, utilizando a palavra-chave var, crie uma variável com o
	identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você
	criou no exercício anterior.

	2. Na função main:
		2.1 Utilize conversão para transformar o tipo de valor da variável "x" e seu tipo
		subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
		2.2 Demostre o valor de "y"
		2.3 Demonstre o tipo de "y"
*/

type customInt int

var x customInt = 2004
var y int

func main() {
	fmt.Printf("\n%d\n%T", x, x)

	x = 42
	fmt.Printf("\n%d", x)

	y = int(x)
	fmt.Printf("\n%d\n%T", y, y)
}
