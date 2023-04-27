package model

import (
	"github.com/TskFok/AdminApi/global"
	tool_page "github.com/TskFok/AdminApi/tool/tool-page"
	"github.com/TskFok/AdminApi/utils/logger"
)

type User struct {
	BaseModel
	Id       uint32 `gorm:"column:id;type:INT UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Email    string `gorm:"column:email;type:VARCHAR(255);NOT NULL" json:"email,omitempty"`
	Salt     string `gorm:"column:salt;type:VARCHAR(50);NOT NULL" json:"salt,omitempty"`
	Password string `gorm:"column:password;type:VARCHAR(255);NOT NULL" json:"password,omitempty"`
	Status   int8   `gorm:"column:status;type:TINYINT(1);NOT NULL" json:"status"`
}

func (u *User) Find(info *User) *User {
	db := global.DataBase.Where(info).First(&u)

	if db.Error != nil {
		logger.Error(db.Error.Error())
	}

	return u
}

func (u *User) List(page, size int) map[string]interface{} {
	db := global.DataBase.Select("id,email,created_at,updated_at,status").
		Offset((page - 1) * size).Limit(size).Order("id desc")

	uList := make([]User, size)
	return tool_page.Paginate(db, &uList)
}

func (*User) Add(user *User) uint32 {
	db := global.DataBase.Create(user)

	if db.Error != nil {
		logger.Error(db.Error.Error())

		return 0
	}

	return user.Id
}

func (*User) Update(user *User, update map[string]interface{}) bool {
	db := global.DataBase.Model(&User{}).Where(user).Updates(update)

	if db.Error != nil {
		logger.Error(db.Error.Error())
		return false
	}

	return true
}
