package main

import (
	connections "08-rest-methods/helpers"
	"fmt"
)

func main() {
	resp, err := connections.GET()
	if err != nil {
		fmt.Printf("Error GET: %s\n", err)
	}
	fmt.Printf("Response:GET: %s\n", resp)

	resp, err = connections.POST()
	if err != nil {
		fmt.Printf("Error POST: %s\n", err)
	}
	fmt.Printf("Response:POST: %s\n", resp)

	resp, err = connections.PUT()
	if err != nil {
		fmt.Printf("Error PUT: %s\n", err)
	}
	fmt.Printf("Response:PUT: %s\n", resp)

	resp, err = connections.DELETE()
	if err != nil {
		fmt.Printf("Error DELETE: %s\n", err)
	}
	fmt.Printf("Response:DELETE: %s\n", resp)

}
