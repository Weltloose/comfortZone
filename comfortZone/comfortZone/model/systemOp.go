package model

import (
	"fmt"
	"os"
)

// CheckLogined checks whether cookie is available
func CheckLogined(uuid string) (bool, error) {
	exist, err := checkAuth(uuid)
	if err != nil {
		return false, fmt.Errorf("check login, %v", err)
	}
	return exist, nil
}

// CheckValid checks whether account username and passwd is matched
func CheckValid(username, passwd string) (bool, string) {
	return validUser(username, passwd), createAuthCookie(username, passwd)
}

// CheckExist checks whether username exist
func CheckExist(username string) bool {
	return checkUserExist(username)
}

// WHen register user also create private doc
func RegisterUser(username, passwd string) bool {
	return createUserInfo(username, passwd) && createPrivateDocs(username)
}

// GetUserWithAuth returns username according to
func GetUsernameWithAuth(uuid string) string {
	username, _, err := readByAuth(uuid)
	if err != nil {
		return ""
	}
	return username
}

func ChangePassword(username, passwd string) bool {
	return updateUserPasswd(username, passwd)
}

func getPrivateDockPath(username string) string {
	return "static/private/" + username
}

func createPrivateDocs(username string) bool {
	filePath := getPrivateDockPath(username)
	_, err := os.Create(filePath)
	if err != nil {
		return false
	}
	return true
}
