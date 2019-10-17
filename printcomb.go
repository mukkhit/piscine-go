package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for l := i + 1; l <= '9'; l++ {
			for k := l + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(l)
				z01.PrintRune(k)
				if i != '7' || l != '8' || k != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')

				}
			}

		}
	}
	z01.PrintRune('\n')
}
