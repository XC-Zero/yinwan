package _interface

// ChineseTabler 带表名和中文表名的模型
//	普适用于 mysql mongo
//	接口名字是传承 gorm 的 Tabler
type ChineseTabler interface {
	// TableCnName 表中文名
	TableCnName() string
	// TableName 表名
	TableName() string
}

// EsTabler 带表名和中文表名的ES模型
type EsTabler interface {
	ChineseTabler
	// Mapping 返回 ES 表结构
	Mapping() map[string]interface{}
	// ToESDoc 转为 ES 的一条记录
	ToESDoc() map[string]interface{}
}
