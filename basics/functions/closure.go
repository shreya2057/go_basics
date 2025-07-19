package functions

import "fmt"

// Closure have ability to remember the variable value even after the function execution is completed.

func closureFunction() {
	closure := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	firstInstance := closure()
	fmt.Println("First instance of closure:", firstInstance()) // Output: 1
	fmt.Println("First instance of closure:", firstInstance()) // Output: 2
	fmt.Println("First instance of closure:", firstInstance()) // Output: 3

	fmt.Println()

	secondInstance := closure()
	fmt.Println("Second instance of closure:", secondInstance()) // Output: 1
	fmt.Println("Second instance of closure:", secondInstance()) // Output: 2
	fmt.Println("Second instance of closure:", secondInstance()) // Output: 3
}
