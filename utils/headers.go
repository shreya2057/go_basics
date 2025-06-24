package utils

import (
	"fmt"
	"strings"
)

func Heading(heading string, number string) {
	var length = len(heading) + 10
	var headingText = fmt.Sprintf("\033[1m**-- %s --**\033[0m", heading)
	fmt.Println(headingText)
	if number == "1" {
		fmt.Println(strings.Repeat("=", length))
	}
	if number == "2" {
		fmt.Println(strings.Repeat("-", length))
	}
}
