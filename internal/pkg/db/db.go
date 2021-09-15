package db

import (
	"im/internal/pkg/logger"
	"sync"

	"gorm.io/gorm"
)

var (
	dbs   map[string]*gorm.DB
	rwMux sync.RWMutex
)

/*
	if you want to adapter other sql driver, please impl interface controller and option
*/
type DBController interface {
	NewDB(DBOption) (*gorm.DB, error)
	SetOption(*gorm.DB, DBOption)
}

type DBOption interface {
	MustEmbedDefaultOption()
}

func GetDB(dbName string, controller DBController, option DBOption) *gorm.DB {
	rwMux.RLock()
	var (
		db  *gorm.DB
		err error
	)

	db, ok := dbs[dbName]

	if !ok {
		rwMux.Lock()
		db, err = controller.NewDB(option)

		if err != nil {
			logger.Panic(err)
		}

		dbs[dbName] = db
		rwMux.Unlock()
	}

	controller.SetOption(db, option)
	rwMux.RUnlock()

	return db
}

type DefaultDBOption struct{}

func (ctrl DefaultDBOption) MustEmbedDefaultOption() {
	logger.Error("must embed default option struct in a customed db option struct")
}
