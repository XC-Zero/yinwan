package mongo_model

type Assemble struct {
	BasicModel   `bson:"inline"`
	BookNameInfo `bson:"-"`
}

func (a Assemble) TableName() string {
	return "assembles"
}
func (a Assemble) TableCnName() string {
	return "组装拆卸"
}
