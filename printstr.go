package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	for _, i := range str {
		z01.PrintRune(i)
	}
}
