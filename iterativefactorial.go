package piscine

func IterativeFactorial(nb int) int {
	arg := 1
	if nb > 0 && nb < 2147483647 {
		for i := 1; i < nb+1; i++ {
			arg = arg * i
		}
		return arg
	}
	return 0
}
