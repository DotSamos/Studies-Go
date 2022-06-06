package main

import "fmt"

/*
	Use var para declarar três variáveis. ELas
	devem ter package-level scope. Não atribua
	valores a estas variáveis. Utilize os seguintes
	identificadores e tipos para esta variáveis:

		1. Identificador "x" deverá ter o tipo int
		2. Identificador "y" deverá ter o tipo string
		3. Identificador "z" deverá ter o tipo bool

	Na função main:

		1. Demonstre os valores de cada identificador
		2. O copilador atribuiu valores para estas variáveis. Como
		estes valores se chamam?
			-	Valores zero
*/

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v, %v, %v", x, y, z)
}
