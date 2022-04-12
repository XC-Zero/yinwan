package tools

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// SliceConvert 数组类型转换
//  此处用了大量反射，高并发情况下将导致性能瓶颈
func SliceConvert(slice interface{}, newSlice interface{}) (interface{}, error) {
	ot, nt := reflect.TypeOf(slice), reflect.TypeOf(newSlice)
	ote, nte := ot.Elem(), nt.Elem()

	if ot.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("Slice called with non-slice value of type %T", ot))
	}
	if nt.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("Slice called with non-slice type of type %T", nt))
	}

	var v = reflect.ValueOf(slice)
	var l = v.Len()
	var dv = reflect.MakeSlice(nt, 0, l)

	if !ote.ConvertibleTo(nte) {
		return nil, errors.New(fmt.Sprintf("Type %T can not convert to type %T", ote, nte))
	}
	for i := 0; i < l; i++ {
		dv = reflect.Append(dv, v.Index(i).Convert(nte))
	}

	return dv.Interface(), nil
}
