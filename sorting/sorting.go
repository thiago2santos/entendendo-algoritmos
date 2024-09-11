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

func OrdenarRapido[T Number](array []T) (arrayConcatenado []T) {
	if len(array) < 2 {
		return array
	} else {
		iPivo := len(array) / 2
		vPivo := array[iPivo]
		menores := make([]T, 0)
		maiores := make([]T, 0)
		for _, e := range array {
			if e < vPivo {
				menores = append(menores, e)
			} else if e > vPivo {
				maiores = append(maiores, e)
			}
		}
		return slices.Concat(append(OrdenarRapido(menores), vPivo), OrdenarRapido(maiores))
	}

}
