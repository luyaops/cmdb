package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	"github.com/luyaops/cmdb/models"
)

var DB *gorm.DB

func Initialize() (*gorm.DB, error) {
	//dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", "root:123yyqwe@tcp(127.0.0.1)/cmdb?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
