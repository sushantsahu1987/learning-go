package main

import "fmt"

func main() {
	fmt.Println("Slices")

	slices1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(slices1)
	fmt.Println(len(slices1))
	fmt.Println(cap(slices1))

	slices2 := slices1[1:]

	fmt.Println(slices2)
	fmt.Println(len(slices2))
	fmt.Println(cap(slices2))

	index := 3
	slices3 := append(slices1[:index], slices1[index+1:]...)
	fmt.Println(slices3)
	fmt.Println(len(slices3))
	fmt.Println(cap(slices3))
	slices1 = slices3
	fmt.Println(slices1)

}
