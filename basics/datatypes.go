package datatypes

import "fmt"

func DataTypes() {
	integer()
}

func integer() {
	var integer int = 10
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 142535664757857444
	fmt.Println("**---Integer--**")
	fmt.Println("int", integer)
	fmt.Println("int8", integer8)
	fmt.Println("int16", integer16)
	fmt.Println("int32", integer32)
	fmt.Println("int64", integer64)
}
