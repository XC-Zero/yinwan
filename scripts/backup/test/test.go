package main

import "github.com/XC-Zero/yinwan/pkg/utils/file_plus"

func main() {
	err := file_plus.Zip(`./`, `./x.zip`)
	if err != nil {
		panic(err)
	}
}
