package main

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

func OpenExcelFromMultiPart(multi *multipart.FileHeader) (*excelize.File, error) {
	f, err := multi.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return nil, err
	}
	return excel, nil
}

var mapping map[int]string

func init() {
	mapping = make(map[int]string, 26)
	for i := 1; i < 27; i++ {
		mapping[i] = string(('A') + int32(i-1))
	}
}

// ColIndexToColString
//	数字下标转换为 A B C D...下标
//	index 按照正常编程逻辑从0开始
func ColIndexToColString(index int) string {
	return math_plus.TenToAnyWithMapping(index+1, 26, mapping)
}

// CellIndexToCellString
//	单元格数字下标转换为 A4 B7 ... 这种下标
//	index 按照正常编程逻辑从0开始
func CellIndexToCellString(colIndex, rowIndex int) string {
	return ColIndexToColString(colIndex) + strconv.Itoa(rowIndex+1)
}

func NewRootCell() *cellHeader {
	return &cellHeader{
		colIndex: 0,
		rowIndex: 0,
		width:    0,
		height:   1,
		content:  nil,
		father:   nil,
		children: make([]*cellHeader, 0),
	}
}

type cellHeader struct {
	colIndex int
	rowIndex int
	width    int
	height   int
	content  interface{}
	father   *cellHeader
	children []*cellHeader
}

func (c *cellHeader) AddChild(newChild *cellHeader, isXAxis bool) {
	newChild.father = c

	nowWidth, nowHeight := 0, 0
	for i := 0; i < len(c.children); i++ {
		nowWidth += c.children[i].width
		nowHeight += c.children[i].height
	}

	if isXAxis {
		newChild.colIndex = c.colIndex + nowWidth
		newChild.rowIndex = c.rowIndex + 1
		c.children = append(c.children, newChild)
		c.GrowWidth()
	} else {
		newChild.colIndex = c.colIndex + 1
		newChild.rowIndex = c.rowIndex + nowHeight
		c.children = append(c.children, newChild)
		c.GrowHeight()
	}
	spew.Dump(newChild)
	log.Println("_________________________________________________")
}

// GrowWidth
//	此处考虑是 子节点+多少  父节点也加多少
//  还有个方案就是 子节点变动,父节点统计求和,这样的话不允许父节点比子节点宽一点
func (c *cellHeader) GrowWidth() {
	var w int
	for i := 0; i < len(c.children); i++ {
		if c.children[i] != nil {
			w += c.children[i].width
		}
		c.width = w
	}
	if c.width == 0 {
		c.width = 1
	}
	if c.father != nil {
		c.father.GrowWidth()
	}
}
func (c *cellHeader) GrowHeight() {
	var h int
	for i := 0; i < len(c.children); i++ {
		if c.children[i] != nil {
			h += c.children[i].height
		}
		c.height = h
	}
	if c.height == 0 {
		c.height = 1
	}
	if c.father != nil {
		c.father.GrowHeight()
	}
}

// VerticalAddMap cellHeader 竖着添加 map
func (c *cellHeader) VerticalAddMap(m map[string]interface{}) {
	for key, value := range m {
		objType := reflect.TypeOf(value)
		objValue := reflect.ValueOf(value)
		if objType.Kind() == reflect.Ptr {
			objValue = objValue.Elem()
		}
		switch objValue.Kind() {
		case reflect.Struct:
			continue
		case reflect.Map:
			kid := cellHeader{
				width:    1,
				height:   1,
				content:  key,
				father:   c,
				children: make([]*cellHeader, 0),
			}
			kid.VerticalAddMap(objValue.Interface().(map[string]interface{}))
			c.AddChild(&kid, true)
		case reflect.Slice:
			continue
		case reflect.Array:
			continue
		default:
			c.AddChild(&cellHeader{
				width:    1,
				height:   1,
				content:  key,
				father:   c,
				children: make([]*cellHeader, 0),
			}, true)
		}
	}
}

func StructToMap(i interface{}) (map[string]interface{}, error) {
	if reflect.TypeOf(i).Kind() != reflect.Struct {
		return nil, errors.Errorf("Unsupported %t type! ", i)
	}
	valMap := make(map[string]interface{}, 0)
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshal, &valMap)
	if err != nil {
		return nil, err
	}
	return valMap, nil
}

func (c *cellHeader) WriteIntoSheet(file *excelize.File, sheetName string) error {

	if c.father != nil {
		log.Printf("index is %s,value is %s", CellIndexToCellString(c.colIndex, c.rowIndex), c.content)
		err := file.SetCellValue(sheetName, CellIndexToCellString(c.colIndex, c.rowIndex), c.content)
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(c.children); i++ {
		err := c.children[i].WriteIntoSheet(file, sheetName)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (c *cellHeader) AddStruct(model interface{}, tag string) {
	if reflect.TypeOf(model).Kind() != reflect.Struct {
		return
	}
	objV := reflect.ValueOf(model)
	objT := reflect.TypeOf(model)
	for i := 0; i < objV.NumField(); i++ {
		tagV := objT.Field(i).Tag.Get(tag)
		if tagV == "" {
			tagV = objT.Field(i).Name
		}
		switch reflect.TypeOf(objV.Field(i).Interface()).Kind() {
		case reflect.Struct:
			kid := cellHeader{
				width:    1,
				height:   1,
				content:  tagV,
				father:   c,
				children: make([]*cellHeader, 0),
			}
			kid.AddStruct(objV.Field(i).Interface(), tag)
			c.AddChild(&kid, true)
		case reflect.Map:
			kid := cellHeader{
				width:    1,
				height:   1,
				content:  tagV,
				father:   c,
				children: make([]*cellHeader, 0),
			}
			kid.VerticalAddMap(objV.Field(i).Interface().(map[string]interface{}))
			c.AddChild(&kid, true)
		case reflect.Slice:
			continue
		case reflect.Array:
			continue
		default:
			c.AddChild(&cellHeader{
				width:    1,
				height:   1,
				content:  tagV,
				father:   c,
				children: make([]*cellHeader, 0),
			}, true)
		}

	}
}
