package database

import (
	"github.com/jinzhu/gorm"
	"github.com/mvrilo/app-poc/pkg/config"

	// database adapters
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

func New() (*Database, error) {
	db, err := gorm.Open(config.DatabaseAdapter(), config.DatabaseURI())
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
