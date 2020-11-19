package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digital(67.779887777 * 100))
}

//RoundFloatNumbers save float64 with 2 digitals
func RoundFloatNumbers(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
