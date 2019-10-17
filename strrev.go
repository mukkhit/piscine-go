package piscine

func StrRev(s string) string { 
	var d string
	for _, i := range s {
        d = string(i) + d
	}
      return d
}
