package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power >= 0 {

		for i := 1; i < (power + 1); i++ {
			result = nb * result
		}
	} else {
		result = 0
	}
	return result
}
