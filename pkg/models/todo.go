package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Uuid        *string        `gorm:"primaryKey" json:"uuid" binding:"required"`
	OwnerUuid   *string        `gorm:"index" json:"ownerUuid" binding:"required"`
	State       *string        `json:"state" binding:"required"`
	Title       *string        `json:"title" binding:"required"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
