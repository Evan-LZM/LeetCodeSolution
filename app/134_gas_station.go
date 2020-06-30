package app

func canCompleteCircuit(gas []int, cost []int) int {
	total, sum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		sum += gas[i] - cost[i]
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
