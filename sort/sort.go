package sort

import "fmt"

func ordenarPorSelecao[T Number](array []T) (arrayOrdenado []T) {
	for i := 0; i < len(array); {
		menor := GetMin(array)
		arrayOrdenado = append(arrayOrdenado, array[menor])
		antes := array[:menor]
		depois := array[menor:]
		fmt.Printf("Antes: %v, Depois: %v", antes, depois)
	}
	return
}
