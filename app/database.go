package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"levante/orm"
	"levante/util"
	"log"
	"os"
)

var db *gorm.DB

func InitDatabase(config *AppConfig) *gorm.DB{
	var err error
	dbConnURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Schema)
	db, err = gorm.Open("mysql", dbConnURL)
	if err != nil {
		panic("failed to connect database")
	}
	setDBLogger(config)
	db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&orm.Post{})
	db.AutoMigrate(&orm.LinkGroup{})
	db.AutoMigrate(&orm.Link{})
	db.AutoMigrate(&orm.User{})
	return db
}

func setDBLogger(config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.ApplicationLog.File)
	if !util.CheckIsExistPath(logPath) {
		panic("logPath:" + logPath + " is not exist! bootstrap.go must execute setLogger() before setDatabase()")
	}
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	db.SetLogger(log.New(f, "", 0))
}
