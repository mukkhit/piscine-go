package piscine

func TrimAtoi(s string) int {
	final := 0
	sign := 1
	for _, com := range s {
		if com >= '0' && com <= '9' {
			final = final*10 + int(com-48)
		} else if com == '-' && final == 0 {
			sign = -1
		}
	}
	return final * sign
}
