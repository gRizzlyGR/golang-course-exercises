package abs

// Abs computes the absolute value
func Abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
