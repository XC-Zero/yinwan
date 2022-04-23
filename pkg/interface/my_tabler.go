package _interface

type ChineseTabler interface {
	TableCnName() string
	TableName() string
}

type EsTabler interface {
	ChineseTabler
	Mapping() string
}
