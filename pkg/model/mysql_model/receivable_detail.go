package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ReceivableDetail struct {
	BasicModel
	BookNameInfo
	ReceivableID           int     `gorm:"type:int" json:"receivable_id,omitempty" form:"receivable_id,omitempty"`
	ReceivableAmount       *string `gorm:"type:varchar(50)" json:"receivable_amount,omitempty" form:"receivable_amount,omitempty"`
	ReceivableOperatorID   *string `gorm:"type:varchar(50)" json:"receivable_operator_id,omitempty" form:"receivable_operator_id,omitempty"`
	ReceivableOperatorName *string `gorm:"type:varchar(50)" json:"receivable_operator_name,omitempty" form:"receivable_operator_name,omitempty"`
	Remark                 *string `gorm:"type:varchar(200)" json:"remark,omitempty"  form:"remark"`
}

func (p ReceivableDetail) TableCnName() string {
	return "应付详情"
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
	receivable.RecID = &p.ReceivableID
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
	receivable.ReceivableActualAmount = &actualAmount
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Updates(&receivable).Where("rec_id = ?", receivable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应付记录失败!")
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
