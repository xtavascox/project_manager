package model

import "github.com/google/uuid"

type Project struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey" validate:"required"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Description string    `json:"description" gorm:"not null" validate:"required"`
	CreatorId   uuid.UUID `json:"creator_id" gorm:"not null" validate:"required"`
	CreatedAt   int64     `json:"created_at" gorm:"not null" `
	ModifiedAt  int64     `json:"modified_at" gorm:"not null"`
}
