package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteController struct {
}

type SqliteOption struct {
	DefaultDBOption

	DBPath          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdletime time.Duration
}

func (controller SqliteController) NewDB(option DBOption) (*gorm.DB, error) {
	sqliteOption := option.(SqliteOption)

	db, err := gorm.Open(sqlite.Open(sqliteOption.DBPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (controller SqliteController) SetOption(db *gorm.DB, option DBOption) {
	sqliteDB, _ := db.DB()
	sqliteOption := option.(SqliteOption)

	sqliteDB.SetMaxIdleConns(sqliteOption.MaxIdleConns)
	sqliteDB.SetMaxOpenConns(sqliteOption.MaxOpenConns)
	sqliteDB.SetConnMaxLifetime(sqliteOption.ConnMaxLifetime * time.Second)
	sqliteDB.SetConnMaxIdleTime(sqliteOption.ConnMaxLifetime * time.Second)
}
