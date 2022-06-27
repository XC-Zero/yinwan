package main

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	m "github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	es "github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

var esIndexes []_interface.EsTabler

func init() {
	esIndexes = append(esIndexes,
		&es.Material{},
		&es.Commodity{},
		&es.Payable{},
		&es.Receivable{},
		&es.FixedAsset{},
		&es.Provider{},
		&es.Customer{},
		&m.FinanceCredential{},
		&m.StockOutRecord{},
		&m.StockInRecord{},
		&m.Purchase{},
		&m.Transaction{},
		&m.Return{},
	)

}

func GenerateESIndex() {
	for _, index := range esIndexes {
		err := client.CreateIndex(index)
		if err != nil {
			panic(err)
		}
	}
}

func DropESIndex() {
	for _, index := range esIndexes {
		client.DeleteIndex(index)

	}
}
