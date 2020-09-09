package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username    string
	Email       string `gorm:"unique;index;"`
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ActivatedAt *time.Time
}

func (user User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New())
	if err != nil {
		return err
	}
	return nil
}
