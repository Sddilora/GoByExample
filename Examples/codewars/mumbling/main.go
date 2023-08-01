// This time no story, no theory. The examples below show you how to write function accum:

// Examples:
// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
package main

import (
	"fmt"
	"strings"
)

func main() {
	print(Accum("codewars"))
}

func Accum(s string) string {

	sArr := make([]string, len(s))

	for i, v := range s {
		sArr[i] = string(v)
	}

	fmt.Println(sArr) // [c o d e w a r s]
	strResult := ""

	for i := 0; i < len(sArr); i++ {
		ltrUp := strings.ToUpper(sArr[i])
		ltrLow := strings.Repeat(strings.ToLower(sArr[i]), i)
		ltrUpLow := ltrUp + ltrLow
		if i == len(sArr)-1 {
			strResult += ltrUpLow
		} else {
			strResult += ltrUpLow + "-"
		}
	}

	return strResult
}
