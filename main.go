package main

import (
	"fmt"
)

func main() {
	// 1. Basic: Loop and If Statement demonstration
	fmt.Println("=== Basic Demonstration ===")
	for i := 1; i <= 3; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
	fmt.Println()

	// 2. Structure (Struct) demonstration:
	fmt.Println("=== Structure Demonstration ===")
	// Define a Person struct
	type Person struct {
		Name string
		Age  int
	}
	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", person)
	fmt.Println()

	// 3. Slice demonstration with open-ended extension
	fmt.Println("=== Slice Demonstration ===")
	// Create an initial slice of integers
	numbers := []int{1, 2, 3}
	fmt.Printf("Initial slice: %v (len: %d, cap: %d)\n", numbers, len(numbers), cap(numbers))
	// Append elements to the slice, extending it beyond its original capacity
	numbers = append(numbers, 4, 5, 6)
	fmt.Printf("Extended slice: %v (len: %d, cap: %d)\n", numbers, len(numbers), cap(numbers))
	fmt.Println()

	// 4. Array demonstration:
	fmt.Println("=== Array Demonstration ===")
	// Declare an array of 5 integers
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array: %v\n", arr)
}
