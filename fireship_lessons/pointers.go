package main

import "fmt"

// The default value for pointers is `nil` (like `null` in JS)

func main() {
	var address *int  // In this case, the * indicates the pointer type (int).
	number := 42      // int
	address = &number // & is the address operator (address stores the memory address of number)
	value := *address // In this case, the * dereferences the value to which the `address` points.

	fmt.Println("address: \n", address)
	fmt.Println("value: \n", value)
}
