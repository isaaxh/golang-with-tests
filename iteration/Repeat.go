package iteration

import (
	"strings"
)

// Repeat takes a string and an integer and returns the string repeated the integer times.
func Repeat(c string, r int) string {
	var repeated strings.Builder
	for i := 0; i < r; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}
