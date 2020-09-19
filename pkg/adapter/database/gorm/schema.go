package gorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          string `gorm:"type:uuid;primary_key;"`
	Email       string `gorm:"unique;index;"`
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ActivatedAt *time.Time
}

func (user User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}
