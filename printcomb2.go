package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for k := '0'; k <= '9'; k++ {
			for l := '0'; l <= '9'; l++ {
				for j := '0'; j <= '9'; j++ {
					if i == l && k < l || i < l {
					z01.PrintRune(i)
					z01.PrintRune(k)
					z01.PrintRune(' ')
					z01.PrintRune(l)
					z01.PrintRune(j)
					if i != '9' || k != '8' || l != '9' || j != '9' {

						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
}
	z01.PrintRune('\n')
}
