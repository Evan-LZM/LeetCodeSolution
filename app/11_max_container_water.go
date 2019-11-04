package app

func MaxArea(height []int) int {
	maxarea := 0
	i, j := 0, len(height)-1
	for {
		if i == j {
			break
		}
		h := height[i] //store the highest
		if height[i] > height[j] {
			h = height[j]
		}
		area := (j - i) * h
		if area > maxarea {
			maxarea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxarea
}
