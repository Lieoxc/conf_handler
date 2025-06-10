package model

import (
	"fmt"

	"gorm.io/gorm"
)

const StatusExpect = "expect"
const StatusOk = "normal"

type Urls struct {
	ID         uint   `gorm:"primary_key;auto_increment" json:"id"`
	Urls       string `gorm:"type:varchar(128);not null" json:"urls"`
	Status     string `gorm:"type:varchar(50);not null" json:"status"`
	NodeCount  int    `gorm:"type:INT(11);not null" json:"node_count"`
	ExpiryDate string `gorm:"type:varchar(50);not null" json:"expiry_date"`
}

// GetCate 查询分类列表
func GetUrls() ([]Urls, error) {
	var urls []Urls
	err = db.Find(&urls).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("get failed")
		return []Urls{}, err
	}
	return urls, nil
}
func UpdateUrlInfo(url, status string, count int) error {
	if err := db.Model(&Urls{}).Where(" urls = ?", url).Updates(Urls{Status: status, NodeCount: count}).Error; err != nil {
		return err
	}
	return nil
}
