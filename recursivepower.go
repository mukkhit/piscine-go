package piscine

func RecursivePower(nb int, power int) int {
	{
		if nb >= 0 && nb < 2147483647 {
			if nb == 1 || nb == 0 {
				return 1
			}
			if nb > 1 {
				return RecursivePower(nb, power-1) * nb
			}
		}
		return 0
	}

}
