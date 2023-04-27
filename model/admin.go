package model

import "github.com/TskFok/AdminApi/global"

type Admin struct {
	BaseModel
	Id       uint32 `gorm:"column:id;type:INT UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Email    string `gorm:"column:email;type:VARCHAR(255);NOT NULL" json:"email,omitempty"`
	Salt     string `gorm:"column:salt;type:VARCHAR(50);NOT NULL" json:"salt,omitempty"`
	Password string `gorm:"column:password;type:VARCHAR(255);NOT NULL" json:"password,omitempty"`
	Status   int8   `gorm:"column:status;type:TINYINT(1);NOT NULL" json:"status,omitempty"`
}

func (*Admin) Find(condition map[string]interface{}) (u *Admin) {
	global.DataBase.Where(condition).First(&u)

	return
}
