package app

import "sort"

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]city)
	for _, value := range tickets {
		m[value[0]] = append(m[value[0]], city{
			name:    value[1],
			visited: false,
		})
	}
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i].name < v[j].name
		})
	}
	result := []string{"JFK"}
	helperItinerary(m, "JFK", len(tickets), &result)
	return result
}

func helperItinerary(m map[string][]city, name string, count int, res *[]string) bool {
	if count == 0 {
		return true
	}
	list := m[name]
	for i := 0; i < len(list); i++ {
		to := list[i]
		if to.visited {
			continue
		}
		list[i].visited = true
		*res = append(*res, to.name)
		if helperItinerary(m, to.name, count-1, res) {
			return true
		}
		*res = (*res)[:len(*res)-1]
		list[i].visited = false
	}
	return false
}
