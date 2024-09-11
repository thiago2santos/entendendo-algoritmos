package sort

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func GetMin[T Number](arr []T) (indice int) {
	indice = 0
	min := arr[0]
	for i, e := range arr {
		if e < min {
			min = e
			indice = i
		}
	}
	return
}
