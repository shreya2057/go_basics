package complexNumber

import (
	"fmt"
	"go_lang/utils/headers"
)

func Complex() {
	var complex complex64 = 200983889986789863837479777777777777777i
	var complex2 complex128 = 128e+306i

	headers.Heading("Complex", "")
	fmt.Println("complex64(64 bits)", complex)
	fmt.Println("complex128(128 bits)", complex2)
	fmt.Println()
}
