package main

// Max returns the greater of two integer values
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of two integer values
func Min(x int, y int) int {
	return -Max(-x, -y)
}
