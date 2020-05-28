package app

func Trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var result int
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	lmax := 0
	rmax := 0
	for i := 0; i < len(height); i++ {
		//store left
		leftMax[i] = max(lmax, height[i])
		lmax = max(lmax, height[i])

		//store right
		rightMax[len(height)-i-1] = max(rmax, height[len(height)-i-1])
		rmax = max(rmax, height[len(height)-i-1])
	}
	for index, value := range height {
		result += minTrap(leftMax[index], rightMax[index]) - value
	}
	return result
}

func minTrap(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxTrap(a, b int) int {
	if a > b {
		return a
	}
	return b
}
