package main

import (
	"fmt"
)

var i int
var soma int
var prod int

func inserir(num *[10]int) {
	fmt.Printf("Insira 10 números:\n")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d - ", i)
		fmt.Scanf("%d\n", &num[i])
	}
}

func mostrar(num *[10]int) {
	fmt.Printf("Array:\n")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", num[i])
	}
}

func media(num *[10]int) float32 {
	soma = 0
	for i = 0; i < 10; i++ {
		soma = soma + num[i]
	}
	return float32(soma) / 10
}

func somatoria(num *[10]int) int {
	soma = 0
	for i = 0; i < 10; i++ {
		soma = soma + num[i]
	}
	return soma
}

func produtoria(num *[10]int) int {
	prod = 1
	for i = 0; i < 10; i++ {
		prod = prod * num[i]
	}
	return prod
}

func main() {
	fmt.Printf("MENU:\n")
	var vector [10]int
	for {
		var opcao int
		fmt.Printf("Escolha uma opção: ")
		fmt.Printf("\n[1] inserir números")
		fmt.Printf("\n[2] calcular a média")
		fmt.Printf("\n[3] calcular a somatória")
		fmt.Printf("\n[4] calcular a produtória")
		fmt.Printf("\n[5] mostrar array")
		fmt.Printf("\n[6] sair")
		fmt.Printf("\n---> ")
		fmt.Scanf("%d", &opcao)
		switch opcao {
		case 1:
			inserir(&vector)
		case 2:
			fmt.Printf("A media dos valores é %f\n", media(&vector))
		case 3:
			fmt.Printf("A somatória dos valores é %d\n", somatoria(&vector))
		case 4:
			fmt.Printf("A produtória dos valores é %d\n", produtoria(&vector))
		case 5:
			mostrar(&vector)
		case 6:
			return
		}
	}
}
