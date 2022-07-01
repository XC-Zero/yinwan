package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/xuri/excelize/v2"
	"log"
)

type test struct {
	A bool
	B struct {
		C string `json:"cccccc"`
	}
	D string
	E struct {
		E2 int
		E3 float64
		E4 struct {
			EE5 string
		}
	}
	F rune
}

func main() {
	excel := excelize.NewFile()
	excel.NewSheet("test01x")
	a := test{}
	c := NewRootCell()
	c.AddStruct(a, "json")
	log.Println("___________________________________________________")
	spew.Dump(c)

	err := c.WriteIntoSheet(excel, "test01x")
	if err != nil {
		panic(err)
	}

	//err := excel.SetSheetRow("test01x", "A3", a)
	//if err != nil {
	//	panic(err)
	//}
	//err = excel.SetRowOutlineLevel("test01x", 5, 2)
	//if err != nil {
	//	panic(err)
	//}
	err = excel.SaveAs("wtf.xlsx")
	if err != nil {
		panic(err)
	}

}
