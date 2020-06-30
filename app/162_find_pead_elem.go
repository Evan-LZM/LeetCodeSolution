package app

func findPeakElement(n []int) int {
	if len(n) <= 1 {
		return 0
	}
	if n[0] > n[1] {
		return 0
	}
	if n[len(n)-1] > n[len(n)-2] {
		return len(n) - 1
	}
	for i := 1; i < len(n)-1; i++ {
		if n[i] > n[i-1] && n[i] > n[i+1] {
			return i
		}
	}
	return -1
}
