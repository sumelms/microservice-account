package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email       string    `gorm:"unique;index;"`
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ActivatedAt *time.Time
}

// BeforeCreate executes a hook before create the database entry
func (user User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}
