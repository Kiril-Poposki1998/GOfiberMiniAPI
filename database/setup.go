package database

import (
	"runtime"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"
	"github.com/gofiber/storage/redis/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(mysql:3306)/person?charset=utf8mb4&parseTime=True&loc=Local"
var Redis *redis.Storage

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
	Database.AutoMigrate(&model.User{})
	return nil
}

func ConnectRedis() error {
	Redis = redis.New(redis.Config{
		Host:      "redis",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
	if Redis == nil {
		panic("Could not connect to redis")
	}
	return nil
}
