package mongo_model

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

// SaleInvoice 销售订单
type SaleInvoice struct {
	mysql_model.BasicModel
	SaleInvoiceOwnerID   int                    `json:"sale_invoice_owner_id"`
	SaleInvoiceOwnerName string                 `json:"sale_invoice_owner_name"`
	SaleInvoiceContent   map[string]interface{} `json:"sale_invoice_content"`
	SaleAmount           string                 `json:"sale_amount"`
	CustomerID           int                    `json:"customer_id"`
	CustomerName         string                 `json:"customer_name"`
}

func (s SaleInvoice) TableCnName() string {
	return "销售订单"
}
func (s SaleInvoice) TableName() string {
	return "sale_invoices"
}

func (s SaleInvoice) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"customer_name": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"sale_invoice_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"created_at": mapping{
					"type": "text",
				},
			},
		},
	}
	return ma
}

// ImportInvoice 导入单据
// todo
func (s *SaleInvoice) ImportInvoice(invoiceType _const.InvoiceType) (error, inv _interface.Invoice) {
	return nil, nil
}

// ExportInvoice 导出单据
// todo

func (s *SaleInvoice) ExportInvoice() error {

	return nil
}

// TransferToCredential 单据转凭证
// 重点关照
// todo 重点关照
func (s *SaleInvoice) TransferToCredential(creType _const.CredentialType) _interface.Credential {
	return nil
}
