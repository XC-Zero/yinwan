package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	m "github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	es "github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func SelectMaterial(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")

	query := elastic.NewMultiMatchQuery(searchContent, "rec_id^1000", "remark^2", "material_name^10")

	op := common.SelectESTemplateOptions{
		TableModel:  &es.Material{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectCommodity(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")

	query := elastic.NewMultiMatchQuery(searchContent, "rec_id^1000", "remark^2", "commodity_name^10")

	op := common.SelectESTemplateOptions{
		TableModel:  &es.Commodity{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectPayable(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")
	query := elastic.NewMultiMatchQuery(searchContent,
		"rec_id^1000",
		"payable_amount^1000",
		"remark^2",
		"payable_enterprise_address^3",
		"payable_contact^5",
		"payable_enterprise^10")

	op := common.SelectESTemplateOptions{
		TableModel:  &es.Payable{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectReceivable(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")
	query := elastic.NewMultiMatchQuery(searchContent,
		"rec_id^1000",
		"receivable_amount^1000",
		"remark^2",
		"receivable_enterprise_address^3",
		"receivable_contact^5",
		"receivable_enterprise^10")

	op := common.SelectESTemplateOptions{
		TableModel:  &es.Receivable{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectFixedAsset(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")
	query := elastic.NewMultiMatchQuery(searchContent,
		"rec_id^1000",
		"fixed_asset_name.keyword^500",
		"remark^2",
		"fixed_asset_name^3",
		"fixed_asset_amount^5")

	op := common.SelectESTemplateOptions{
		TableModel:  &es.FixedAsset{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectStockInRecord(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")
	query := elastic.NewMultiMatchQuery(searchContent,
		"rec_id^1000",
		"stock_in_owner.keyword^500",
		"remark^2",
		"stock_in_owner^3",
		"stock_in_content^5")

	op := common.SelectESTemplateOptions{
		TableModel:  &m.StockInRecord{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}

func SelectStockOutRecord(ctx *gin.Context) {
	searchContent := ctx.PostForm("search_content")
	query := elastic.NewMultiMatchQuery(searchContent,
		"rec_id^1000",
		"stock_out_owner.keyword^500",
		"remark^2",
		"stock_out_owner^3",
		"stock_out_content^5")

	op := common.SelectESTemplateOptions{
		TableModel:  &m.StockOutRecord{},
		Query:       query,
		ResHookFunc: nil,
	}
	common.SelectESTableContentWithCountTemplate(ctx, op)
	return
}
