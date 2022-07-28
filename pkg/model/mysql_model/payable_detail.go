package mysql_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PayableDetail struct {
	BasicModel
	BookNameInfo
	PayableID           *int    `gorm:"type:int" json:"payable_id,omitempty" form:"payable_id,omitempty" cn:"关联应付编号"`
	PayableDate         *string `gorm:"type:varchar(50)" json:"payable_date,omitempty" form:"payable_date,omitempty" cn:"付款时间"`
	PayableContent      *string `gorm:"type:varchar(200)" json:"payable_content,omitempty" form:"payable_content,omitempty" cn:"付款内容摘要"`
	PayableAmount       *string `gorm:"type:varchar(50)" json:"payable_amount,omitempty" form:"payable_amount,omitempty" cn:"付款金额"`
	PayableOperatorID   *int    `json:"payable_operator_id,omitempty" form:"payable_operator_id,omitempty" cn:"操作人编号"`
	PayableOperatorName *string `gorm:"type:varchar(50)" json:"payable_operator_name,omitempty" form:"payable_operator_name,omitempty" cn:"操作人名称"`
	PayablePicUrlList   *string `gorm:"type:mediumtext" form:"payable_pic_url_list,omitempty" json:"payable_pic_url_list,omitempty" cn:"相关单据图片"`
	Remark              *string `gorm:"type:varchar(200)" json:"remark,omitempty"  form:"remark" cn:"备注"`
}

func (p PayableDetail) TableCnName() string {
	return "应收详情"
}
func (p PayableDetail) TableName() string {
	return "payable_details"
}

// AfterCreate todo 修改实际付款，剩余付款，付款状态
func (p *PayableDetail) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	p.BookNameID = bk.StorageName
	p.BookName = bk.BookName

	var payable Payable
	payable.RecID = p.PayableID
	var amountList []string
	err := bk.MysqlClient.WithContext(tx.Statement.Context).
		Raw("select payable_amount from payable_details where deleted_at is null and payable_id = ?",
			p.PayableID,
		).Scan(&amountList).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询应付记录详情失败!")
		return err
	}
	logger.Info(fmt.Sprintf("amountList is %+v", amountList))
	num := math_plus.Zero
	for _, amount := range amountList {
		fraction, err := math_plus.NewFromString(amount)
		if err != nil {
			logger.Error(errors.WithStack(err), "转化失败!")
			continue
		}
		num = num.Add(fraction)
	}
	if p.PayableAmount != nil {
		f, err := math_plus.NewFromString(*p.PayableAmount)
		if err != nil {
			return err
		}
		num = num.Add(f)
	}
	actualAmount := num.String()
	//查询数据库中payable的值
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Where("rec_id = ?", payable.RecID).Find(&payable).Error
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
	unpaid := _const.UnPaid
	paid := _const.Paid
	partialPaid := _const.PartialPaid

	switch {
	case num.IsZero():
		payable.PayableStatus = &unpaid
	case sub.IsZero():
		payable.PayableStatus = &paid
	case actualAmount < *payable.PayableTotalAmount:
		payable.PayableStatus = &partialPaid
	}
	err = bk.MysqlClient.WithContext(tx.Statement.Context).Where("rec_id = ?", payable.RecID).Updates(&payable).Error
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
