package headers

import (
	"fmt"
	"strings"
)

func Heading(heading string, number string) {
	var length = len(heading) + 10
	var headingText = fmt.Sprintf("**-- %s --**", heading)
	fmt.Println(headingText)
	if number == "1" {
		fmt.Println(strings.Repeat("=", length))
	}
	if number == "2" {
		fmt.Println(strings.Repeat("-", length))
	}
}
