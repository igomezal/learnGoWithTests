package iteration

import "strings"

func Repeat(character string, repeatedTimes int) string {
	repeated := ""
	for i := 0; i < repeatedTimes; i++ {
		repeated += character
	}
	return repeated
}

func OriginalRepeat(character string, repeatedTimes int) string {
	return strings.Repeat(character, repeatedTimes)
}
