package main

import "fmt"

/*
	O operador curto de declaração é utilizado para **declarar** variáveis, desta forma
	caso uma variável já exista este não pode ser utilizado para reescrever o valor dela,
	é necessário utilizar um operador de atribuição para isso (go complicando a vida :v)

	O lado bom de utilizar o operador curto de atribuição é que ele já faz a tipagem
	automática da variável
*/

var name = "Samoca" // variável declarada no contexto global do pacote, todas as funções aqui dentro tem acesso a ela
// age := 18 /* Apenas dentro do escopo das funções (isso é triste) */

func main() {

	// declaramos a variável message e seguindo a tipagem automática ela já vai ser do tipo string
	message := "Hello world!"

	// Imprimindo a variável name que está no escopo global
	fmt.Printf("\nMeu nome é %s", name)

	name = "Samos" // alterando o valor de uma variável no escopo global
	fmt.Printf("\nMeu nome *correto* é %s", name)

	// Igual o 'sprintf' do PHP
	fmt.Printf("\nValor da variável `message` (%T): %s", message, message)

	// message := "Hello new world!" /* O operador curto de atribuição não pode reescrever o valor de uma variável */

	// para reescrever o valor de message é necessário utilizar o operador de atribuição "="
	message = "Hello new world!"

	fmt.Printf("\nNovo valor da variável `message` (%T): %s", message, message)

	// message foi declarado inicialmente sendo uma string, desta forma ela não pode ser redeclarada como int
	// message = 13012004
	// fmt.Printf("\nNovo valor da variável `message` (%T): %s", message, message)

	// operadores e operandos

	// declarando uma variável do tipo int
	number := 20

	fmt.Printf("\nO número base é %d", number)

	// operações com o número
	// number = number + 2;
	// number = number - 2;
	// number = number / 2;
	// number = number % 2;
	number = number * 2

	fmt.Printf("\nO número 20 multiplicado por 2 é %d", number)
}
