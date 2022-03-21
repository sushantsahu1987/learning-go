package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices")

	slices1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("slice 1 ", slices1)
	fmt.Println("length of slice 1 ", len(slices1))
	fmt.Println("capacity of slice 1 ", cap(slices1))

	slices2 := slices1[1:]

	fmt.Println("")
	fmt.Println("slice 2 ", slices2)
	fmt.Println("length of slice 2 : ", len(slices2))
	fmt.Println("capacity of slice 2 : ", cap(slices2))

	slices3 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	index := 2
	slices3[index] = slices3[len(slices3)-1]
	slices3[len(slices3)-1] = 0
	slices3 = slices3[:len(slices3)-1]
	fmt.Println("delete index 3 from a slice", slices3)

}
