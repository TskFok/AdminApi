package tool_page

import (
	"github.com/TskFok/AdminApi/utils/logger"
	"gorm.io/gorm"
)

func Paginate(tx *gorm.DB, desc any) map[string]interface{} {
	db := tx.Find(desc)

	if db.Error != nil {
		logger.Error(db.Error.Error())
	}

	var count int64
	tx = tx.Count(&count)

	if tx.Error != nil {
		logger.Error(tx.Error.Error())
	}

	res := make(map[string]interface{})
	res["data_list"] = desc
	res["count"] = count

	return res
}
