package main

import (
	"fmt"
	"reflect"
)

func main() {
	// This declares and defines the array
	languages := [9]string{"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust"} // <--- Note: Must include the trailing comma

	// Define the slices (kinda like the slice method on JS arrays)
	classics := languages[0:3]  // or simply languages[3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include index 3, exclude index 7 (just like JS)
	new := languages[7:9]       // or simply languages[7]

	fmt.Println("Classic languages: %v\n", classics)
	fmt.Println("Modern languages: %v\n", modern)
	fmt.Println("New languages: %v\n", new)

	allLangs := languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}
	jsFrameworks := frameworks[0:4:4]
	frameworks = append(frameworks, "Metor")

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("JS frameworks: %v\n", jsFrameworks)
}
