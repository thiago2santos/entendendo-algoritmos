package main

import (
	"fmt"

	"github.com/thiago2santos/entendendo-algoritmos/sorting"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	// primeiroElementoDaLista := List[int]{nil, 1}
	// segundoElemento := List[int]{&primeiroElementoDaLista, 46}
	// fmt.Println(segundoElemento)
	arr := []int{131, 940, 632, 902, 454, 469, 352, 381, 185, 920}
	fmt.Println(arr)
	sorting.OrdenarPorSelecao[int](arr)
}
