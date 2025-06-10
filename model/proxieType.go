package model

import "fmt"

type ProxiesTypes struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	TypeName string `gorm:"type:varchar(128);not null" json:"type_name"`
	Desc     string `gorm:"type:varchar(128);not null" json:"desc"`
}

func GetTypeNameFromID(id string) (string, error) {
	var typeName string
	// 查询对应 server 的 name
	if err := db.Model(&ProxiesTypes{}).Where("id = ?", id).Pluck("type_name", &typeName).Error; err != nil {
		fmt.Println("Failed to query names:", err)
		return "", nil
	}
	return typeName, nil
}
