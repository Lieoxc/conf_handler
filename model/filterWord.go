package model

import "gorm.io/gorm"

type filterWord struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Word string `gorm:"type:varchar(128);not null" json:"word"`
}

func GetWords() ([]filterWord, error) {
	var words []filterWord
	err = db.Find(&words).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return []filterWord{}, err
	}
	return words, nil
}
