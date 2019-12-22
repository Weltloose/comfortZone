package model

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createUserInfo(username, passwd string) bool {
	dbOpt := db.Create(&UserInfo{
		Username: username,
		Passwd:   passwd,
	})
	if dbOpt.Error != nil {
		fmt.Println("db create error ", dbOpt.Error)
		return false
	}
	if dbOpt.RowsAffected == 0 {
		return false
	}
	return true
}

func checkUserExist(username string) bool {
	user := []UserInfo{}
	dbOpt := db.Where("username = ?", username).Find(&user)
	if dbOpt.Error != nil {
		fmt.Println("db where error, ", dbOpt.Error)
		return false
	}
	if len(user) != 1 {
		return false
	}
	return true
}

func validUser(username, passwd string) bool {
	var user []UserInfo
	dbOpt := db.Where("username=? AND passwd=?", username, passwd).Find(&user)
	if dbOpt.Error != nil {
		fmt.Println("db no valid userinfo", dbOpt.Error)
		return false
	}
	if len(user) != 1 {
		return false
	}
	return true
}
