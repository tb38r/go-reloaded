package tgr

func IsPunc(category rune) bool {
	switch category {
	case
		'.',
		',',
		'!',
		':',
		';',
		'?':
		return true
	}
	return false
}
