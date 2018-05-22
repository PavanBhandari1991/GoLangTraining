package utilpackage

func reversetwo(s string) string {
	r := []rune(s)
	i := 0
	j := len(r)-1
	for i, j ; i < len(r)/2 ; i++, j-- {
		r[i], r[j] = r[j], r[i]
	}
	return string(s)
}
