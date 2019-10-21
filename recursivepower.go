package piscine

func RecursivePower(nb int, power int) int {
	result := nb
	if power > 0 {
		result *= RecursivePower(nb, power-1)
	} else if power == 0 {
		result = 1
	} else {
		result = 0
	}
	return result
}
