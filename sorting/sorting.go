package sorting

import (
	"fmt"
	"slices"
)

func OrdenarPorSelecao[T Number](array []T) (arrayOrdenado []T) {
	for i := 0; i < len(array); {
		menor := GetMin(array)
		arrayOrdenado = append(arrayOrdenado, array[menor])
		antes := array[:menor]
		depois := array[menor+1:]
		fmt.Printf("Antes: %v, Depois: %v", antes, depois)
		array = slices.Concat(antes, depois)
	}
	return
}
