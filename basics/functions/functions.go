package functions

import (
	"fmt"
	"go_lang/utils"
)

func funcWithNoParameters() {
	utils.FormatOutput("Functions", "No parameters", "This function does not take any parameters and does not return any values.")
}

func functWithSameParameters(paramter1, paramter2 int) int {
	return1 := paramter1 + paramter2
	return return1
}

func functWithTwoReturns(parameter1, parameter2 int) (int, string) {
	if parameter2 == 0 {
		return 0, "Cannot divide by zero"
	}
	return parameter1 / parameter2, ""
}

func funcWithNamedReturn(parameter1, parameter2 int) (result int, err string) {
	if parameter2 == 0 {
		err = "Cannot divide by zero"
		result = 0
	} else {
		err = "Success"
		result = parameter1 / parameter2
	}
	return
}

func funcWithNameReturnShadow() (return1 int) {
	return1 = 10
	if true {
		return1 := 20
		fmt.Println("Inner return1:", return1) // This will print the inner return1
	}
	fmt.Println("Outer return1:", return1) // This will print the outer return1
	return
}

func Functions() {
	utils.Heading("Functions", "1")
	paramter1 := 20
	paramter2 := 10
	funcWithNoParameters()
	add := functWithSameParameters(paramter1, paramter2)
	utils.FormatOutput("Functions With same type paramters", fmt.Sprintf("paramter1: %d, parameter2: %d", paramter1, paramter2), fmt.Sprintf("result of function: %d", add))
	quotient, err := functWithTwoReturns(paramter1, paramter2)
	if err != "" {
		utils.FormatOutput("Functions with two returns", fmt.Sprintf("parameter1: %d, parameter2: %d", paramter1, paramter2), fmt.Sprintf("error: %s", err))
	} else {
		utils.FormatOutput("Functions with two returns", fmt.Sprintf("parameter1: %d, parameter2: %d", paramter1, paramter2), fmt.Sprintf("result of function: %d", quotient))
	}

	value, err := funcWithNamedReturn(paramter1, paramter2)
	utils.FormatOutput("Functions with named returns", fmt.Sprintf("parameter1: %d, parameter2: %d", paramter1, paramter2), fmt.Sprintf("result: %d, error: %s", value, err))

	value = funcWithNameReturnShadow()
	utils.FormatOutput("Functions with name return shadow", "No parameters", fmt.Sprintf("result: %d", value))
}
