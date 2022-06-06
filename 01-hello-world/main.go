package main // definir a package (desconheço isso, deve ser tipo as namespaces do PHP)

import "fmt" // importamos a biblioteca que controla o input/output do go (baseado em C)

func main() {
	/*
		No go as funções podem ter múltiplos retornos, por isso é possível definir
		mais de uma variável antes de invocar uma função, as variáveis serão preenchidas
		de acordo com a sequência do retorno da função
	*/
	numberOfBytesWritten, error := fmt.Println("Hello world", "Let's Go ^^")
	// esse : antes do = me mata

	// aqui podemos imprimir os retornos da função de print
	fmt.Println(numberOfBytesWritten, error)

	/*
		A política do Go é sempre focar em ter o melhor desempenho possível, e por isso
		não permite declarar variáveis que não são utilizadas em nenhum momento, porém
		divido a sequência de retorno das funções para imprimir, por exemplo, apenas o
		erro é necessário "ignorar" o primeiro retorno que seria o número de bytes que
		foram escritos, para isso ao invés de declarar o número de bytes sendo uma variável
		comum, basta declarar ela como "_"
	*/

	// numberOfBytesWritten, error := fmt.Println("Hello world 2", "Let's Go ^^") /* ERROR */

	// aparentemente ele não gosta de reescrever variáveis... * Eu não sei o conceito ainda nesta linguagem :v *
	_, exception := fmt.Println("Hello world", "Let's Go ^^")

	fmt.Println("Erro durante a execução:", exception)

	// <nil> como retorno ao que aparente é equivalente ao NULL de outras linguagens
}
