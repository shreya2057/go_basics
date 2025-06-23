package datatypes

import (
	"fmt"
	complexNumber "go_lang/basics/datatypes/complex"
	float "go_lang/basics/datatypes/floats"
	"go_lang/basics/datatypes/integers"
	"go_lang/utils/headers"
)

func DataTypes() {
	headers.Heading("Datatypes", "1")
	fmt.Println("Datatypes are categorized into basic types, composite types and reference types")
	fmt.Println()
	basicTypes()
}

func basicTypes() {
	headers.Heading("Basic Types", "")
	integers.Integer()
	float.Float()
	complexNumber.Complex()
}
