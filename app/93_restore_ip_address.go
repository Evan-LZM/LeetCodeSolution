package app

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {

	var result []string
	helperRestoreIpAddress(s, []string{}, &result)
	return []string{}
}

func helperRestoreIpAddress(s string, temp []string, data *[]string) {
	if len(temp) == 4 && len(s) == 0 {
		*data = append(*data, strings.Join(temp, "."))
		return
	}
	if len(temp) == 4 && len(s) != 0 {
		return
	}
	for i := 1; i <= 3; i++ {
		if i > len(s) {
			continue
		}
		num, _ := strconv.Atoi(s[:i])
		if s[:i] != strconv.Itoa(num) || num > 255 {
			continue
		}
		temp = append(temp, s[:i])
		helperRestoreIpAddress(s[i:], temp, data)
		temp = temp[:len(temp)-1]
	}
}
