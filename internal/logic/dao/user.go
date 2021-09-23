package dao

import (
	"fmt"
	"reflect"
	"time"

	"im/internal/pkg/db"
	"im/pkg/common"
	"im/pkg/config"
)

type UserModel struct {
	UserID         uint64    `gorm:"primarykey" model:"user_id"`
	UserName       string    `model:"username"`
	SaltedPassword string    `model:"saltedPassword"`
	Creatime       time.Time `model:"createtime"`
}

const tableName = "user"

var (
	dbName       = config.GetConfig().Logic.DB.DBName
	dbController = db.SqliteController{}
	dbOption     = db.SqliteOption{
		DBPath:          config.GetConfig().Logic.DB.Sqlite.DBPath,
		MaxIdleConns:    config.GetConfig().Logic.DB.Sqlite.MaxIdleConns,
		MaxOpenConns:    config.GetConfig().Logic.DB.Sqlite.MaxOpenConns,
		ConnMaxLifetime: config.GetConfig().Logic.DB.Sqlite.ConnMaxLifetime,
		ConnMaxIdletime: config.GetConfig().Logic.DB.Sqlite.ConnMaxIdletime,
	}
	logicDB = db.GetDB(dbName, dbController, dbOption)
)

func newUserID() uint64 {
	/*
		write uid generate rules here
	*/
	return 1
}

func Create(user *UserModel) error {
	_, err := Read(user.UserID)

	if err == nil {
		return common.ErrUserHasExisted
	}

	if user.UserName == "" {
		return common.ErrInvaildUserName
	}

	if len(common.UnsaltPassword(user.SaltedPassword)) < 8 {
		return common.ErrInvaildPassword
	}

	user.Creatime = time.Now()

	if err = logicDB.Table(tableName).Create(&user).Error; err != nil {
		return err
	}
	/*
		Create will auto generate userid, you can also generate uid byself:

		user.UserID = newUserID()
	*/

	return nil
}

func Read(userID uint64) (*UserModel, error) {
	var user *UserModel
	err := logicDB.Table(tableName).First(user, userID).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func ReadByName(userName string) (*UserModel, error) {
	field, _ := reflect.TypeOf(UserModel{}).FieldByName("UserName")
	tag := field.Tag
	nameTag := tag.Get("modle")
	sqlStmt := fmt.Sprintf("%s=?", nameTag)

	var user *UserModel
	err := logicDB.Table(tableName).Where(sqlStmt).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func Update(userID uint64, newUser UserModel) error {
	user, err := Read(userID)

	if err != nil {
		return err
	}

	if err = logicDB.Model(user).Updates(newUser).Error; err != nil {
		return err
	}

	return nil
}

func UpdateField(userID uint64, fieldName string, fieldValue interface{}) error {
	user, err := Read(userID)

	if err != nil {
		return err
	}

	if err = logicDB.Model(user).Update(fieldName, fieldValue).Error; err != nil {
		return err
	}

	return nil
}

func Delete(userID int) error {
	var user *UserModel

	if err := logicDB.Delete(user, userID).Error; err != nil {
		return err
	}

	return nil
}
