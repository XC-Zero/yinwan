package model

import "time"

// BasicModel borrowed from GORM
type BasicModel struct {
	// Currently ID is set to be VARCHAR(100) in order to
	// support different kind of IDs from various 3rd-party
	// sources
	RecID *int `gorm:"primaryKey;type:int;autoIncrement" json:"id"`

	// Every row should at least mark up the create, update
	// and delete timestamp
	CreatedAt time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
}

// TimeOnlyModel does not have an ID
type TimeOnlyModel struct {
	CreatedAt time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
}
