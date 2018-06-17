package caesar

import (
	"fmt"
	"strings"
)

// Bruteforce tries out all possible values for shift
func Bruteforce(in string) {
	for i := 1; i <= 26; i++ {
		fmt.Println(Apply(in, i))
	}
}

// Apply substitues all runes of in by shift
func Apply(in string, shift int) string {
	var out []rune
	pt := strings.ToLower(in)
	for _, x := range pt {
		if x == ' ' {
			out = append(out, ' ')
			continue
		}

		if x < 'a' || x > 'z' {
			continue
		}

		unilateralShift := (x - 97) + int32(shift)
		if unilateralShift < 0 {
			unilateralShift = 26 + unilateralShift
		}

		r := (unilateralShift % 26) + 97
		out = append(out, r)
	}
	return strings.ToUpper(string(out))
}
