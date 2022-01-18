package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString('\n')
	// please finish the function body here.
	// please define the output here. For example: fmt.Println(input)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		list := strings.Split(string(line), " ")
		a, err := strconv.ParseFloat(list[0], 64)
		if err != nil {
			return
		}
		b, err := strconv.Atoi(list[1])
		if err != nil {
			return
		}
		if a == 0 && b == 0 {
			return
		}

	}

}
