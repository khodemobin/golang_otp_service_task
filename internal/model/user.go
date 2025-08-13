package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey" column:"id"`
	Phone     string `column:"phone" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
