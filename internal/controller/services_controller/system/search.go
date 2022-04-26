package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/es_model"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func SelectMaterial(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")

	query := elastic.NewMultiMatchQuery(searchContent, "rec_id^1000", "remark^2", "material_name^10")

	op := common.SelectESTemplateOptions{
		TableModel:  es_model.Material{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}
