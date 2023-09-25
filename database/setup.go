package database

import (
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(mysql:3306)/person?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&model.Person{})
	return nil
}
