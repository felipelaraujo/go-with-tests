package iteration

import "strings"

func Repeat(text string, target int) (repeatedText string) {
	// manual
	// for i := 0; i < target; i++ {
	// 	repeatedText += text
	// }

	return strings.Repeat(text, target)
}
