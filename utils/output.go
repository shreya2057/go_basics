package utils

import "fmt"

func FormatOutput(method string, input string, output interface{}) {
	fmt.Println()
	fmt.Printf("\033[1m%s\033[0m\n", method)
	fmt.Println("Input:", input, "|", "Output:", output)
	fmt.Println()
}
