package storage

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"gorm.io/gorm"
)

// AddCommodity 新增一类商品
func AddCommodity(db *gorm.DB, commodity model.Commodity) bool {
	return false
}

// UpdateCommodity 修改商品
func UpdateCommodity(db *gorm.DB, commodity model.Commodity) bool {
	//db.Model(&model.Commodity{}).Where().Update("", "")
	return false
}

// DeleteCommodity 删除商品
func DeleteCommodity(db *gorm.DB, commodity model.Commodity) bool {
	return false
}

// SelectCommodity 查找商品
func SelectCommodity(db *gorm.DB, commodity model.Commodity) []model.Commodity {
	return nil
}
