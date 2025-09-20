package main

import (
	"fmt"
)

func main() {

	fmt.Println("Arrays e Slices")
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice4 := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice4)

	slice4 = append(slice4, 18)
	fmt.Println(slice4)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays Internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice5 := make([]float32, 5)
	fmt.Println(slice5)
	slice5 = append(slice5, 10)
	fmt.Println(len(slice5)) // length
	fmt.Println(cap(slice5)) // capacity




}

