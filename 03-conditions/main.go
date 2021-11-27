package main

import "fmt"

func main() {
	fmt.Println("Conditions")

	i := 3

	if i%2 == 0 {
		fmt.Printf("Even %d\n", i)
	} else {
		fmt.Printf("Odd %d\n", i)
	}

	i = 4

	if i%2 == 0 {
		fmt.Printf("Even %d\n", i)
	} else {
		fmt.Printf("Odd %d\n", i)
	}

	i = 9

	switch i {
	case 4:
		fmt.Printf("Case 4\n")

	case 5:
		fmt.Printf("Case 5\n")

	default:
		fmt.Printf("Default\n")
	}

}
