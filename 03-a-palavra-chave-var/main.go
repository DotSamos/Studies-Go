package main

import "fmt"

/*
	Para declarar variáveis globais/no "package-level" se utiliza a keyword "var" fora
	do contexto de uma função
*/
var date = "06/jun/2022 13:45"

func main() {
	// podemos acessar a variável date aqui dentro do escopo da função main
	fmt.Printf("\nData em que este exemplo foi feito: %s", date)

	// declaramos uma variável dentro do escopo da função main
	language := "Go"
	fmt.Printf("\nEstamos utilizando a linguagem '%s'", language)

	// test1()
	test2(language) // invocamos a função só que agora passando a variável language como parâmetro
}

func test1() {
	// fmt.Printf("\n%s é uma linguagem ruim :v", language); /* isto não funciona, pois language foi declarado no contexto da função main */
}

func test2(language string) {
	fmt.Printf("\n%s é uma boa linguagem :)", language)
}
