package gofetch

import (
	"fmt"
	"strings"
)
// for ------------
func PrintWithUnderline(text string) {
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", len(text)))
}