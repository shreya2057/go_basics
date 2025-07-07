package functions

import (
	"fmt"
	"go_lang/utils"
)

func functions(paramter1, parameter2 int, paramter4 string) (int, int) {
	result1 := paramter1 + parameter2
	result2 := len(paramter4)
	return result1, result2
}

func Functions() {
	result1, result2 := functions(10, 20, "Hello World")
	utils.Heading("Functions", "1")
	utils.FormatOutput("Functions", fmt.Sprintf("paramter1: %d, parameter2: %d, paramter3: %s", 10, 20, "Hello world"), fmt.Sprintf("result1: %d, result2: %d", result1, result2))
}
