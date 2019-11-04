package app

import "fmt"

type roman struct {
	num int
	s   string
}

func IntToRoman(num int) string {
	romans := []roman{
		roman{
			num: 1000,
			s:   "M",
		},
		roman{
			num: 900,
			s:   "CM",
		},
		roman{
			num: 500,
			s:   "D",
		},
		roman{
			num: 400,
			s:   "CD",
		},
		roman{
			num: 100,
			s:   "C",
		},
		roman{
			num: 90,
			s:   "XC",
		},
		roman{
			num: 50,
			s:   "L",
		},
		roman{
			num: 40,
			s:   "XL",
		},
		roman{
			num: 10,
			s:   "X",
		},
		roman{
			num: 9,
			s:   "IX",
		},
		roman{
			num: 5,
			s:   "V",
		},
		roman{
			num: 4,
			s:   "IV",
		},
		roman{
			num: 1,
			s:   "I",
		},
	}
	result := ""
	for i := 0; i < len(romans); i++ {
		o := num / romans[i].num
		fmt.Println("o:", o, ",roman_i:", romans[i].s)
		if o > 0 {
			for j := 0; j < o; j++ {
				fmt.Println("j:", j)
				result += romans[i].s
				fmt.Println("roman_s:", romans[i].s)
			}
		}
		num = num - o*romans[i].num
		if num == 0 {
			break
		}
	}
	return result
}
