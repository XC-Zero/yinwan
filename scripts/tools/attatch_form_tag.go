package main

import (
	"io/ioutil"
	"log"
	"os"
)

// 270 330
var fileList map[string]*os.File = make(map[string]*os.File, 0)

func main() {
	err := GetAllFile("D:\\go\\workspace\\yinwan\\pkg\\model")

	if err != nil {
		panic(err)
	}
	for k, v := range fileList {

		var contentStr string
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return
		}
		contentStr = string(all)
		log.Println("************************************************************\n " + k)
		log.Println(contentStr)
		//r, err := regexp.Compile("json:\"(.*)\"")
		//r.FindAll()
		if err != nil {
			return
		}
		for {
			break
		}
		err = os.WriteFile(k, []byte(contentStr), 0644)

	}
}

func GetAllFile(pathname string) (err error) {

	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		panic(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			err := GetAllFile(pathname + "\\" + fi.Name())

			if err != nil {
				panic(err)

				return err
			}
		} else {
			filePath := pathname + "\\" + fi.Name()
			log.Printf("???? is %s", filePath)
			f, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			fileList[filePath] = f
		}
	}
	return
}
