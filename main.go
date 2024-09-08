package main

import "fmt"

func main() {
	helloWorld()
}

func helloWorld() {
	fmt.Println("Hello, World!")
	min := GetMin([]int{5, 4, 3, 2, 1})
	fmt.Println(min)
}

func GetMin(arr []int) (min int) {
	min = arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	return
}
