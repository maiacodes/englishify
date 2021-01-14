package englishify

import (
	"fmt"
	"strings"
)

// FmtNumber will format an input to use Ordinal numbers.
// e.g. int(2) -> "2nd"
func FmtNumber(number interface{}) string {
	input := fmt.Sprint(number)

	if strings.HasSuffix(input, "1") {
		return input + "st"
	}
	if strings.HasSuffix(input, "2") {
		return input + "nd"
	}
	if strings.HasSuffix(input, "3") {
		return input + "rd"
	}

	return input + "th"
}
