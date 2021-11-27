package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	for _, i := range array {
		fmt.Println(i)
	}

	strs := [5]string{
		"a", "b", "c", "d", "e"}

	fmt.Println(strs)

	for _, i := range strs {
		fmt.Println(i)
	}

}
