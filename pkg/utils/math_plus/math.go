package math_plus

import (
	"math"
	"strings"
)

//GreatestCommonDivisor 最大公约数:(辗转相除法)
func GreatestCommonDivisor(x, y int64) int64 {
	x = int64(math.Abs(float64(x)))
	y = int64(math.Abs(float64(y)))

	var tmp int64
	for {
		tmp = x % y
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

//LeastCommonMultiple 最小公倍数:((x*y)/最大公约数)
func LeastCommonMultiple(x, y int64) int64 {
	return (x * y) / GreatestCommonDivisor(x, y)
}

// AnyToTen 任意进制转10进制
func AnyToTen(num string, n int) int {
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

// map根据value找key
func findKey(in string) int {
	result := -1
	for k, v := range mapping {
		if in == v {
			result = k
		}
	}
	return result
}

func TenToAny(num, n int) string {
	return TenToAnyWithMapping(num, n, mapping)
}

var mapping = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// TenToAnyWithMapping 10进制转任意进制
// num 10进制数
// n 几进制
// mapping 映射表
func TenToAnyWithMapping(num, n int, mapping map[int]string) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % n
		remainderString = mapping[remainder]
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}
