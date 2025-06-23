package float

import (
	"fmt"
	"go_lang/utils/headers"
)

func Float() {
	var float float32 = 10.33333333333333333333333333334
	var float2 float64 = 10.3333333333333333333333333333

	headers.Heading("Float", "")
	fmt.Println("float32(Single precision float):", float)
	fmt.Println("float64(Double precision float)", float2)

	fmt.Println()
}
