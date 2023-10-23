package lecture2

import "strings"

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var res strings.Builder

	for num > 0 {
		for i := range value {
			if num >= value[i] {
				res.WriteString(roman[i])
				num -= value[i]
				break

			}
		}

	}

	return res.String()
}
