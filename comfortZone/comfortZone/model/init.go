package model

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var redisClient *redis.Client

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(mysql:3306)/comfort_zone?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("data not opened, ", err)
		return
	}
	db.AutoMigrate(&UserInfo{})
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
