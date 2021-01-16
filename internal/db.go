package internal

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	dbManager = make(map[string]*gorm.DB,0)
)

var (
	DbNotFoundError = errors.New("not found specific databases by key")
	DbAlreadyExistError = errors.New("key already refers to the exist database")
)

func AddDB(key string,db2 *gorm.DB)error{
	if _,ok:=dbManager[key];ok==true{
		return DbAlreadyExistError
	}
	dbManager[key]=db2
	return nil
}

func RemoveDB(key string)bool{
	if _,ok := dbManager[key];ok==true{
		delete(dbManager,key)
		return true
	}
	return false
}

func SetWorkDB(key string)(*gorm.DB,error){
	if db,ok :=dbManager[key];ok==true{
		return db,nil
	}
	return nil,DbNotFoundError
}
