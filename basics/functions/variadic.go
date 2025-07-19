package functions

import (
	"fmt"
	"go_lang/utils"
)

func variadicFunction(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum = sum + number
	}
	return
}

func VariadicFunction() {
	sum := variadicFunction(1, 2, 3, 4, 5)
	utils.FormatOutput("Variadic Function", "Sum of numbers", fmt.Sprintf("The sum of the numbers is: %d", sum))
}
