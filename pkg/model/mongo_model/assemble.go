package mongo_model

type Assemble struct {
	BasicModel   `bson:"inline"`
	BookNameInfo `bson:"-"`
}

func (a Assemble) TableName() string {

}
