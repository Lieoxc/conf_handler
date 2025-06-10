package model

import "gorm.io/gorm"

type CustomRule struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Rule string `gorm:"type:varchar(128);not null" json:"rule"`
}

func GetCustomRule() ([]string, error) {
	var customRules []CustomRule
	err = db.Find(&customRules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return []string{}, err
	}
	rules := make([]string, 0, len(customRules))
	for _, val := range customRules {
		rules = append(rules, val.Rule)
	}
	return rules, nil
}
