package functions

import (
	"fmt"
	"go_lang/utils"
)

// This should never be used as the anonymous function should be used within a scope. It not a best practice.
var anonymous = func() string {
	return "Variable assigned anonymous function"
}

func functWithAnonymousFunction(data string, operation func(string) string) {
	result := operation(data)
	utils.FormatOutput("Anonymous Function as argument", fmt.Sprintf("operation function on: %s", data), result)
}

func Anonymous() {
	function := anonymous()
	utils.FormatOutput("Anonymous function", "No paramters", function)

	func() {
		utils.FormatOutput("IIFE", "No parameters", "This function is executed immediately")
	}()

	output := func() string {
		return "This is anonymous function assigned to a variable"
	}

	utils.FormatOutput("Anonymous Function", "No parameters", output)

	functWithAnonymousFunction("GoLang", func(s string) string {
		return s + " is awesome!"
	})

}
