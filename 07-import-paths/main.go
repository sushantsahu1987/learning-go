package main

import (
	"fmt"

	helpers "07-import-paths/helper"
)

func main() {
	fmt.Println("Import")

	name, age := helpers.Hello()

	// name, age := helper.Hello()
	fmt.Printf("Hello my name is %s and my age is %d\n", name, age)

	// name, age := Helper.Hello()

}
