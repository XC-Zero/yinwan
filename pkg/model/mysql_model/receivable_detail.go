package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ReceivableDetail 收款详情
type ReceivableDetail struct {
	BasicModel
	BookNameInfo
	ReceivableID           *int    `gorm:"type:int" json:"receivable_id,omitempty" form:"receivable_id,omitempty" cn:"关联应收编号"`
	ReceivableDate         *string `gorm:"type:varchar(50)" json:"receivable_date,omitempty" form:"receivable_date,omitempty" cn:"收款日期"`
	ReceivableContent      *string `gorm:"type:varchar(200)" json:"receivable_content,omitempty" form:"receivable_content,omitempty" cn:"收款款内容摘要"`
	ReceivableAmount       *string `gorm:"type:varchar(50)" json:"receivable_amount,omitempty" form:"receivable_amount,omitempty" cn:"收款金额"`
	ReceivableOperatorID   *string `gorm:"type:varchar(50)" json:"receivable_operator_id,omitempty" form:"receivable_operator_id,omitempty" cn:"操作人编号"`
	ReceivableOperatorName *string `gorm:"type:varchar(50)" json:"receivable_operator_name,omitempty" form:"receivable_operator_name,omitempty" cn:"操作人名称"`
	ReceivablePicUrlList   *string `gorm:"type:mediumtext" form:"receivable_pic_url_list,omitempty" json:"receivable_pic_url_list,omitempty" cn:"相关单据图片"`
	Remark                 *string `gorm:"type:varchar(200)" json:"remark,omitempty"  form:"remark" cn:"备注"`
}

func (p ReceivableDetail) TableCnName() string {
	return "收款详情"
}
func (p ReceivableDetail) TableName() string {
	return "receivable_details"
}

func (p *ReceivableDetail) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	p.BookNameID = bk.StorageName
	p.BookName = bk.BookName

	var receivable Receivable
	receivable.RecID = p.ReceivableID
	var amountList []string
	err := bk.MysqlClient.WithContext(tx.Statement.Context).
		Raw("select ? from ?  where ? = ? and ? = ?",
			"receivable_amount", p.TableName(),
			"deleted_at", "is null",
			"receivable_id", p.ReceivableID,
		).Scan(&amountList).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询应付记录详情失败!")
		return err
	}

	err = bk.MysqlClient.WithContext(tx.Statement.Context).Select(&receivable).Where("rec_id = ?", receivable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "应付记录查询失败!")
		return err
	}

	totalAmount, err := math_plus.NewFromString(*receivable.ReceivableTotalAmount)
	if err != nil {
		logger.Error(errors.WithStack(err), "转化失败!")
	}

	num := math_plus.Zero
	for _, amount := range amountList {
		fraction, err := math_plus.NewFromString(amount)
		if err != nil {
			logger.Error(errors.WithStack(err), "转化失败!")
			continue
		}
		num = num.Add(fraction)
	}
	actualAmount := num.String()

	unrec := _const.Unfinished
	rec := _const.Finished
	par := _const.Partial
	sub := totalAmount.Sub(num)
	debtAmount := sub.String()
	receivable.ReceivableActualAmount = &actualAmount
	receivable.ReceivableDebtAmount = &debtAmount
	switch {
	case num.IsZero():
		receivable.ReceivableStatus = &unrec
	case sub.IsZero():
		receivable.ReceivableStatus = &rec
	case actualAmount < *receivable.ReceivableTotalAmount:
		receivable.ReceivableStatus = &par
	}
	receivable.ReceivableActualAmount = &actualAmount
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Updates(&receivable).Where("rec_id = ?", receivable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应收记录失败!")
		return err
	}
	return nil
}

func (p *ReceivableDetail) AfterUpdate(tx *gorm.DB) error {
	return p.AfterCreate(tx)
}

func (p *ReceivableDetail) AfterDelete(tx *gorm.DB) error {
	return p.AfterCreate(tx)
}
