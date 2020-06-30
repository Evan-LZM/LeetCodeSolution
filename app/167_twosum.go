package app

func twoSum2(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if (numbers[i] + numbers[j]) == target {
			return []int{i + 1, j + 1}
		}
		if (numbers[i] + numbers[j]) < target {
			i++
		} else if (numbers[i] + numbers[j]) > target {
			j--
		}
	}
	return nil
}

//another solution
func twoSum21(numbers []int, target int) []int {
	maps := make(map[int]int)
	for k, v := range numbers {
		if m, ok := maps[target-v]; ok {
			return []int{m + 1, k + 1}
		}
		maps[v] = k
	}
	return nil
}
