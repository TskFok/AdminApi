package model

import (
	"fmt"
	"github.com/TskFok/AdminApi/global"
	tool_page "github.com/TskFok/AdminApi/tool/tool-page"
	"github.com/TskFok/AdminApi/utils/logger"
)

type CorpusVicuna struct {
	BaseModel
	Id     uint32 `gorm:"column:id;type:INT UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Corpus string `gorm:"column:corpus;type:LONGTEXT;NOT NULL" json:"corpus,omitempty"`
	Data   string `gorm:"column:data;type:LONGTEXT;NOT NULL" json:"data,omitempty"`
	Pid    uint32 `gorm:"column:pid;type:INT;NOT NULL" json:"pid,omitempty"`
}

func (c *CorpusVicuna) Detail(condition map[string]interface{}) []CorpusVicuna {
	list := make([]CorpusVicuna, 100)

	db := global.DataBase.Select("id,created_at,updated_at,corpus").Where(condition).Order("id desc").Find(&list)

	if db.Error != nil {
		fmt.Println(db.Error.Error())
	}

	return list
}

func (*CorpusVicuna) List(condition map[string]interface{}, page, size int) map[string]interface{} {
	db := global.DataBase.Select("id,created_at,updated_at,corpus").Where(condition).Offset((page - 1) * size).Limit(size).Order("id desc")

	list := make([]CorpusVicuna, size)

	return tool_page.Paginate(db, &list)
}

func (*CorpusVicuna) Add(corpus *CorpusVicuna) uint32 {
	db := global.DataBase.Create(corpus)

	if db.Error != nil {
		logger.Error(db.Error.Error())
	}

	return corpus.Id
}

func (*CorpusVicuna) One(d, c *CorpusVicuna) {
	db := global.DataBase.Where(d).First(&c)

	if db.Error != nil {
		fmt.Println(db.Error.Error())
	}

	return
}

func (*CorpusVicuna) Update(corpus *CorpusVicuna, update map[string]interface{}) bool {
	db := global.DataBase.Model(&CorpusVicuna{}).Where(corpus).Updates(update)

	if db.Error != nil {
		logger.Error(db.Error.Error())

		return false
	}

	return true
}

func (c *CorpusVicuna) Delete() bool {
	db := global.DataBase.Unscoped().Delete(c)

	if db.Error != nil {
		logger.Error(db.Error.Error())

		return false
	}

	return true
}
