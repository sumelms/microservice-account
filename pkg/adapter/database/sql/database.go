package sql

import (
	"database/sql"
	"fmt"

	"github.com/sumelms/microservice-user/pkg/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect(cfg *config.Database) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password)

	db, err := sql.Open(cfg.Driver, connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
