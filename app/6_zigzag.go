package app

import (
	"fmt"
)

func Convert(s string, numRows int) {
	if numRows == 1 {
		fmt.Println(s)
	}
	result := ""
	cyclen := 2*numRows - 2
	n := len(s)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cyclen {
			result = result + string(s[j+i])
			if i != 0 && i != (numRows-1) && j+cyclen-i < n {
				result = result + string(s[j+cyclen-i])
			}
		}
	}
	fmt.Println(result)
}
