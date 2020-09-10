package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sumelms/sumelms/user/pkg/config"
)

func Connect(cfg *config.Database) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DB, cfg.User, cfg.Password)

	db, err := gorm.Open(cfg.Driver, connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
