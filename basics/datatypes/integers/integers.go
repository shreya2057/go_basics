package integers

import (
	"fmt"
	"go_lang/utils/headers"
)

func Integer() {
	var integer int = 10
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 142535664757857444

	var uInteger uint = 2562333
	var uInteger8 uint8 = 254
	var uInteger16 uint16 = 65534
	var uInteger32 uint32 = 1000000000
	var uInteger64 uint64 = 10000000000000000000

	headers.Heading("Integers", "")
	fmt.Println("int(32 or 64 bits):", integer)
	fmt.Println("int8(8-bit integer):", integer8)
	fmt.Println("int16(16-bit integer):", integer16)
	fmt.Println("int32(32-bit integer):", integer32)
	fmt.Println("int64(64-bit integer):", integer64)

	fmt.Println()

	headers.Heading("Unsigned Integer", "")
	fmt.Println("uint(32 or 64 bits):", uInteger)
	fmt.Println("uint8(8-bit unsigned integer):", uInteger8)
	fmt.Println("uint16(16-bit unsigned integer):", uInteger16)
	fmt.Println("uint32(32-bit unsigned integer):", uInteger32)
	fmt.Println("uint64(64-bit unsigned integer):", uInteger64)
}
