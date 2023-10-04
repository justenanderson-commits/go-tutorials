package main

import "fmt"

func main() {
	var address *int  // declares an int pointer
	number := 42      // int
	address = &number // addres stores the memory address of number
	value := *address // dereferencing the value

	fmt.Println("address: \n", address)
	fmt.Println("value: \n", value)
}
