package main

import (
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"gorm.io/gorm"
)

func GenerateTrigger(db *gorm.DB, tables []interface{}) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range tables {
			trigger, ok := tables[i].(_interface.Trigger)
			if ok {
				for _, s := range trigger.Trigger() {
					err := tx.Exec(s).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
