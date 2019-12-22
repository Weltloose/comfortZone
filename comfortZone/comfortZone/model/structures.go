package model

type UserInfo struct {
	ID       int64  `gorm:"AUTO_INCREMENT;primary_key"`
	Username string `gorm:"type:varchar(100); not null; unique"`
	Passwd   string `gorm:"type:varchar(100); not null"`
}

type Comments struct {
	Com []Comment `json:"comments"`
}

type Comment struct {
	CreateTime string `json:"createTime"`
	Username   string `json:"username"`
	Content    string `json:"content"`
}
