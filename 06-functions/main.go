package main

import "fmt"

func Hello() string {
	return "My name is Sushant"
}

func DoSomething() (string, int) {
	return "sushant", 34
}

func main() {
	fmt.Println("Functions")

	sentence := Hello()
	name, age := DoSomething()

	fmt.Println(sentence)
	fmt.Printf("%s is %d years old\n", name, age)

}
