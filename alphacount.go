package piscine

func AlphaCount(str string) int {
	runes := []rune(str)

	count := 0
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= 'a' && runes[i] <= 'z' {
			count++
		}
	}
	return count
}
