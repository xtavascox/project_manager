package model

import "github.com/google/uuid"

type Task struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey" validate:"required"`
	Title        string    `json:"title" gorm:"not null" validate:"required"`
	Description  string    `json:"description" gorm:"not null" validate:"required"`
	CreatorId    uuid.UUID `json:"creator_id" gorm:"not null" validate:"required"`
	AssignedToId uuid.UUID `json:"assigned_to_id" gorm:"not null" validate:"required"`
	ProjectId    uuid.UUID `json:"project_id" gorm:"not null" validate:"required"`
	StatusId     uuid.UUID `json:"status_id" gorm:"not null" validate:"required"`
	StartDate    int64     `json:"start_date" gorm:"not null" validate:"required"`
	EndDate      int64     `json:"end_date" gorm:"not null" validate:"required"`
	CreatedAt    int64     `json:"created_at" gorm:"not null" validate:"required"`
	ModifiedAt   int64     `json:"modified_at" gorm:"not null" validate:"required"`
}
