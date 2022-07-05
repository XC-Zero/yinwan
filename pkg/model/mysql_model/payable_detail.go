package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PayableDetail struct {
	BasicModel
	BookNameInfo
	PayableID           int     `gorm:"type:int" json:"receivable_id,omitempty" form:"receivable_id,omitempty" cn:"关联应付编号"`
	PayableDate         *string `gorm:"type:varchar(50)" json:"receivable_date,omitempty" form:"receivable_date,omitempty" cn:"付款时间"`
	PayableContent      *string `gorm:"type:varchar(200)" json:"receivable_content,omitempty" form:"receivable_content,omitempty" cn:"付款内容摘要"`
	PayableAmount       *string `gorm:"type:varchar(50)" json:"receivable_amount,omitempty" form:"receivable_amount,omitempty" cn:"付款金额"`
	PayableOperatorID   *string `gorm:"type:varchar(50)" json:"receivable_operator_id,omitempty" form:"receivable_operator_id,omitempty" cn:"操作人编号"`
	PayableOperatorName *string `gorm:"type:varchar(50)" json:"receivable_operator_name,omitempty" form:"receivable_operator_name,omitempty" cn:"操作人名称"`
	PayablePicUrlList   *string `gorm:"type:mediumtext" form:"payable_pic_url_list" json:"payable_pic_url_list" cn:"相关单据图片"`
	Remark              *string `gorm:"type:varchar(200)" json:"remark,omitempty"  form:"remark" cn:"备注"`
}

func (p PayableDetail) TableCnName() string {
	return "应收详情"
}
func (p PayableDetail) TableName() string {
	return "receivable_details"
}

// todo 修改实际付款，剩余付款，付款状态
func (p *PayableDetail) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	p.BookNameID = bk.StorageName
	p.BookName = bk.BookName

	var payable Payable
	payable.RecID = &p.PayableID
	var amountList []string
	err := bk.MysqlClient.WithContext(tx.Statement.Context).
		Raw("select ? from ?  where ? = ? and ? = ?",
			"receivable_amount", p.TableName(),
			"deleted_at", "is null",
			"receivable_id", p.PayableID,
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
	//查询数据库中payable的值
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Select(&payable).Where("rec_id = ?", payable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "应付记录查询失败!")
		return err
	}
	totalAmount, err := math_plus.NewFromString(*payable.PayableTotalAmount)
	if err != nil {
		logger.Error(errors.WithStack(err), "转化失败!")
	}
	sub := totalAmount.Sub(num)
	debtAmount := sub.String()
	payable.PayableActualAmount = &actualAmount
	payable.PayableDebtAmount = &debtAmount
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Updates(&payable).Where("rec_id = ?", payable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应付记录失败!")
		return err
	}
	return nil
}

func (p *PayableDetail) AfterUpdate(tx *gorm.DB) error {
	return p.AfterCreate(tx)
}

func (p *PayableDetail) AfterDelete(tx *gorm.DB) error {
	return p.AfterCreate(tx)
}
