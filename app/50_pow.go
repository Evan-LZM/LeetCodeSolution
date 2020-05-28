package app

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	t := MyPow(x, n/2) //speed: n/2>n-1
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}
