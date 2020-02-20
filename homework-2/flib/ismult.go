package flib

// IsMultOfN проверяет целое число x на кратность натуральному mult
func IsMultOfN(x int64, mult int64) bool {
	if x%mult == 0 {
		return true
	}
	return false
}

// IsMultOf2 проверяет целое число x на четность
func IsMultOf2(x int64) bool {
	return IsMultOfN(x, 2)
}

// IsMultOf3 проверяет целое число x на кратность трем
func IsMultOf3(x int64) bool {
	return IsMultOfN(x, 3)
}

