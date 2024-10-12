package utils

import (
	"strings"
	"unicode"
)

func CapitalizeFirstLetter(s string) string {
    if len(s) == 0 {
        return s
    }
    return string(unicode.ToUpper(rune(s[0]))) + strings.ToLower(s[1:])
}
