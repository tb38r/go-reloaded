package tgr

func BasicAtoi2(s string) int {
	srune := []rune(s)
	x := 0

	for i := 0; i < len(srune); i++ {
		if srune[i] >= '0' && srune[i] <= '9' {
			x = x * 10
			x = x + int(srune[i]) - 48
		} else {
			return 0
		}
	}
	return x
}
