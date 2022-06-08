package iteration

import "strings"

// Repeat the character send by parameter any times passed by parameters
func Repeat(character string, times int) string {
	return strings.Repeat(character, times)
}