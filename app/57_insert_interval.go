package app

func Insert(intervals [][]int, newInterval []int) [][]int {
	data := [][]int{}
	joining := false
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			data = append(data, intervals[i])
			continue
		} else if intervals[i][0] > newInterval[1] {
			if !joining {
				joining = true
				data = append(data, newInterval)
			}
			data = append(data, intervals[i])
		} else {
			newInterval[0] = minInsert(newInterval[0], intervals[i][0])
			newInterval[1] = maxInsert(newInterval[1], intervals[i][1])
		}
	}
	if !joining {
		joining = true
		data = append(data, newInterval)
	}
	return data
}

func minInsert(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxInsert(x, y int) int {
	if x > y {
		return x
	}
	return y
}
