// Go does not use a while loop. Instead it uses a for loop to accomplish the same thing.

package main

import "fmt"

func main() {
	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println("Count is: ", count)
}
