package datatypes

import (
	"fmt"
	"go_lang/utils"
)

func DataTypes() {
	utils.Heading("Datatypes", "1")
	fmt.Println("Datatypes are categorized into basic types, composite types and reference types")
	fmt.Println()
	basicTypes()
}

func basicTypes() {
	utils.Heading("Basic Types", "")
	Integer()
	Float()
	Complex()
	String()
}
