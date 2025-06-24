package utils

import (
	"fmt"
	"strings"
)

func Heading(heading string, number string) {
	var length = len(heading) + 10

	// "\033[1m ... \033[0m" bolds text in terminal
	var headingText = fmt.Sprintf("\033[1m**-- %s --**\033[0m", heading)
	fmt.Println(headingText)
	if number == "1" {
		fmt.Println(strings.Repeat("=", length))
	}
	if number == "2" {
		fmt.Println(strings.Repeat("-", length))
	}
}
