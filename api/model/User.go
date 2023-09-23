package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey"`
	FirstName  string    `json:"first_name" validate:"required"`
	LastName   string    `json:"last_name" validate:"required"`
	Email      string    `json:"email"  gorm:"unique" validate:"required"`
	Password   []byte    `json:"password" validate:"required"`
	UserName   string    `json:"user_name"  gorm:"unique" validate:"required"`
	Birthdate  int64     `json:"birthdate" validate:"required"`
	CratedAt   int64     `json:"crated_at" validate:"required"`
	ModifiedAt int64     `json:"modified_at"`
}

type UserDto struct {
	Id        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Birthdate int64     `json:"birthdate"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
