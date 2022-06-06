package main

import "fmt"

/*
	Termos:
		Declaração: Quando você declara
		Inicialização: O primeiro valor que é definido
		Atribuição: Reescrever o valor antigo por um novo

	Quando você declara uma variável sem um valor "padrão" ela vai vir com o valor "zero", este
	valor vai variar de tipo para tipo o qual definiu na variável.

	Os zeros seguindo seus tipos:
		int: 0
		float: 0.0
		boolean: false
		string: ""
		pointers, functions, interfaces, slices, channels, maps: nil
*/

var example float64 // devido a saída ele vai ser impresso como apenas 0, mas o valor real é 0.0

func main() {
	fmt.Println(example)
}
