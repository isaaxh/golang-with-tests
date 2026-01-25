package iteration

import (
	"strings"
)

func Repeat(c string, r int) string {
	var repeated strings.Builder
	for i := 0; i < r; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}
