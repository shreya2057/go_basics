package pointers

import (
	"go_lang/utils"
)

func Pointers() {
	var a int = 20
	var ptr *int = &a
	var nilPtr *int
	var pointerValue int = *ptr
	countPrt := new(int)
	*countPrt = 10
	utils.Heading("Pointers", "1")
	utils.FormatOutput("Pointer to a variable", "ptr", pointerValue)
	utils.FormatOutput("nil value pointer", "nilPtr", nilPtr)
	utils.FormatOutput("new() function", "countPrt", *countPrt)
}
