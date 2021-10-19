package tgr

func IsPunc(category string) bool {
	switch category {
	case
		".",
		",",
		"!",
		":",
		";",
		"?":
		return true
	}
	return false
}
