package main

import "fmt"

/*
	Utilizando a solução do exercício anterior:

	1. Em package-level scope, atribua os seguintes valores as
	variáveis:
		1.1 para "x" atribua 42
		1.2 para "y" atribua "James Bond"
		1.3 para "z" atribua true

	2. Na função main:
		2.1 Use fmt.Sprintf para atribuir todos esses valores a
		uma única variável. Faça essa atribuição de tipo string
		a uma variável de nome "s" utilizando o operador curto
		de declaração
		2.2 Demonstre a variável "s"

*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
