package model

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
)

// SaleInvoice 销售订单
type SaleInvoice struct {
	BasicModel
	SaleInvoiceOwnerID   int    `json:"sale_invoice_owner_id"`
	SaleInvoiceOwnerName string `json:"sale_invoice_owner_name"`
	CustomerID           int    `json:"customer_id"`
	CustomerName         string `json:"customer_name"`
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