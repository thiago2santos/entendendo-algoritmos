package sorting

import (
	"slices"
)

func OrdenarPorSelecao[T Number](array []T) (arrayOrdenado []T) {
	for i := 0; i < len(array); {
		menor := GetMin(array)
		arrayOrdenado = append(arrayOrdenado, array[menor])
		array = slices.Concat(array[:menor], array[menor+1:])
	}
	return
}
