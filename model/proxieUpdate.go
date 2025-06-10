package model

import (
	"gorm.io/gorm"
)

type SubRaw struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`       // 对应 MySQL 的 `id` 字段，主键且自增
	RawUrl  string `gorm:"type:varchar(128);not null" json:"rawUrl"` // 对应 MySQL 的 `raw_data` 字段，类型为 TEXT，非空，字符集为 utf8mb4_unicode_ci
	RawData string `gorm:"type:text;not null" json:"rawData"`        // 对应 MySQL 的 `raw_data` 字段，类型为 TEXT，非空，字符集为 utf8mb4_unicode_ci
}

func GetSubRaw(id string) ([]SubRaw, error) {
	var datas []SubRaw

	// 使用 Where 方法根据 id 过滤查询结果
	err := db.Where("id = ?", id).Find(&datas).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return []SubRaw{}, err
	}

	return datas, nil
}
