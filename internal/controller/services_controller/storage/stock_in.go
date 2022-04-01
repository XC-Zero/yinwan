package storage

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateStockIn(ctx *gin.Context) {
	temp := model.StockInRecord{}

	bookName := ctx.PostForm("book_name")
	err := ctx.ShouldBind(&temp)
	if err != nil || bookName == "" {
		return
	}
	client.BookNameMap[bookName].MongoDBClient.Collection(temp.TableName())
}

func SelectStockInRecord(ctx *gin.Context) {

}
