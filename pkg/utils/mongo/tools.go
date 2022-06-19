package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type OperatorSymbol string

const (
	EQUAL              OperatorSymbol = "$eq"
	NOT_EQUAL          OperatorSymbol = "$ne"
	EXISTS             OperatorSymbol = "$exists"
	LIKE               OperatorSymbol = "$regex"
	IN                 OperatorSymbol = "$in"
	NOT_IN             OperatorSymbol = "$nin"
	GREATER_THEN       OperatorSymbol = "$gt"
	GREATER_THEN_EQUAL OperatorSymbol = "$gte"
	LESS_THAN          OperatorSymbol = "$lt"
	LESS_THAN_EQUAL    OperatorSymbol = "$lte"
)

//var NullTime time.Time
//
//func init() {
//	NullTime, _ = time.Parse("2006-01-02", "1970-01-11")
//}

func TransMysqlOperatorSymbol(symbol OperatorSymbol, column string, value interface{}) bson.E {
	//switch symbol {
	//case EQUAL:
	//	return bson.E{Key: column, Value: value}
	//case LIKE:
	//	return bson.E{Key: column, Value: bson.E{Key: string(LIKE), Value: value}}
	//
	//case IN:
	//	return bson.E{Key: column,Value: bson.E{Key: string()}}
	//case GREATER_THEN:
	//case GREATER_THEN_EQUAL:
	//case LESS_THAN:
	//case LESS_THAN_EQUAL:
	//
	//}
	return bson.E{Key: column, Value: bson.M{string(symbol): value}}
}
