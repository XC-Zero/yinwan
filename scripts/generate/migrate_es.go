package main

import (
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/es_model"
)

var esIndexes []_interface.EsTabler

func init() {
	esIndexes = append(esIndexes,
		es_model.Material{},
		es_model.Commodity{},
		es_model.Payable{},
		es_model.Receivable{},
		es_model.StockOutRecord{},
		es_model.StockInRecord{},
		es_model.FixedAsset{},
	)
}
