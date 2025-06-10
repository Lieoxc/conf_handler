package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SubInfo struct {
	ID         uint            `gorm:"primary_key;auto_increment" json:"id"`
	ConfigType string          `gorm:"type:varchar(50);not null" json:"config_type"`
	Group      string          `gorm:"type:varchar(128);not null" json:"group"`
	Name       string          `gorm:"type:varchar(128);not null" json:"name"`
	Proxies    json.RawMessage `gorm:"type:json;not null" json:"proxies"`
	Url        string          `gorm:"type:TEXT;not null" json:"url"`
}

// GetSubProxies 查询该订阅 所对应订阅地址
func GetSubProxies(token string) ([]Proxies, error) {
	var subInfo SubInfo
	err = db.Where("url LIKE ?", "%"+token+"%").
		Order("id DESC").
		First(&subInfo).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []Proxies{}, nil
		}
		fmt.Println("Error:", err)
		return []Proxies{}, err
	}
	// Parse the proxies JSON field into a string array
	var proxies []Proxies
	if subInfo.ConfigType == "proxies" {
		// subInfo.Proxies => ["zz.dytannel.com:🇰🇷 韩国A丨Oracle丨流媒体", "zz.dytannel.com:🇨🇦 加拿大丨多伦多"]
		prosixStr := []string{}
		err = json.Unmarshal(subInfo.Proxies, &prosixStr)
		if err != nil {
			fmt.Println("Error parsing proxies JSON:", err)
			return []Proxies{}, nil
		}
		for _, valStr := range prosixStr {
			infos := strings.Split(valStr, ":")
			server := infos[0]
			innerName := infos[1]
			proxie, err := GetServersFromServerAndInnerName(server, innerName)
			if err != nil {
				fmt.Println("GetServersFromServerAndInnerName err:", err)
				continue
			}
			proxies = append(proxies, proxie)
		}
	} else if subInfo.ConfigType == "group" {
		typeName, err := GetTypeNameFromID(subInfo.Group)
		if err != nil {
			fmt.Println("GetTypeNameFromID Error:", err)
			return []Proxies{}, nil
		}
		proxies, err = GetServersFromType(typeName)
		if err != nil {
			fmt.Println("GetServersFromType Error:", err)
			return []Proxies{}, nil
		}
	}

	return proxies, nil
}
