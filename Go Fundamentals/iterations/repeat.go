package iteration

import "strings"

// const (
// 	repeatCount = 5
// )

// func Repeat(character string) (repeated string) {
// 	for i := 0; i < repeatCount; i++ {
// 		repeated = character + repeated
// 	}
// 	return
// }

// strings are immutable so every concatenation involves
// copying memory to accomodate for new string
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder // library minimizes memory copying
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character) // method is used to concatenate strings
	}
	return repeated.String()

}
