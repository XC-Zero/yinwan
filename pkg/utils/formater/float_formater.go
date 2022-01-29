package formater

import (
	"fmt"
	"strings"
)

// NumberFormat 输出带逗号且保留俩位的数字
func NumberFormat(num float64) string {
	str := fmt.Sprintf("%.2f", num)
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".")
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".")
}
