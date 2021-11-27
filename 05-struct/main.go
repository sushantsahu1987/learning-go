package main

import "fmt"

func main() {
	fmt.Println("Structures")

	type Person struct {
		Name   string
		Age    int
		Gender string
	}

	persons := []Person{
		{
			Name:   "Sushant",
			Age:    34,
			Gender: "M",
		},
		{
			Name:   "Khushboo",
			Age:    31,
			Gender: "F",
		},
	}

	fmt.Printf("%v", persons)

}
