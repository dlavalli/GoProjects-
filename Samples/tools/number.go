package tools

func Max(a, b int) int {
	// No support for ternary operator in go
	// res := (a > b ? a : b)

	if (a >= b) {
		return a
	}
	return b
}

func Min(a, b int) int {
	if (a <= b) {
		return a
	}
	return b
}

func MaxAreaCheck(width, height, limit int) bool {
	// Shorthand notation allows assignment and condition on the same line
	if area := width * height; area < limit {
		return true
	}
	return false
}

// Version using a switch statement with
func MinAreaCheck(width, height, limit int) bool {
	area := width * height
	switch {
	case area >= limit:
		return true
	default:
		return false
	}
}


