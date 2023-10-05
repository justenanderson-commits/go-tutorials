package main

import "fmt"

func main() {
	c := make(chan int) // This creates a channel to pass integers
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // This starts the goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from the channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // At this point, all goroutines have finished
}

func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // This sends the value back to main
}
