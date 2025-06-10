package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Proxies struct {
	ID         uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(128);not null" json:"name"`
	InnerName  string `gorm:"type:varchar(128);not null" json:"innerName"`
	Server     string `gorm:"type:varchar(128);not null" json:"server"`
	Conf       string `gorm:"type:TEXT;not null" json:"conf"`
	ProxieType string `gorm:"type:varchar(128);not null" json:"proxie_type"`
	FromSub    string `gorm:"type:varchar(128);not null" json:"from_sub"`
}

// CreateArt 新增文章
func CreateProxie(data Proxies) int {
	fmt.Println(data.Name)
	err := db.Create(&data).Error
	if err != nil {
		return -1 // 500
	}
	return 0
}

func UpdateProxie(node Proxies) error {
	// 开始事务
	tx := db.Begin()

	// 查找 server == "asd" 的记录并更新 conf 字段
	if err := tx.Model(&Proxies{}).Where("server = ? and inner_name = ?", node.Server, node.InnerName).Updates(Proxies{Conf: node.Conf, FromSub: node.FromSub}).Error; err != nil {
		tx.Rollback() // 更新出错，回滚事务
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
func CheckExist(addr, innerName string) (bool, error) {
	var count int64
	// 查询 server 是否存在
	if err := db.Model(&Proxies{}).Where("server = ? and inner_name = ?", addr, innerName).Count(&count).Error; err != nil {
		return false, err // 如果有错误，返回错误
	}
	// 如果 count 大于 0，说明存在记录
	if count == 0 {
		fmt.Println("************ count", count, "  ", addr, "  ", innerName)
	}
	return count > 0, nil
}

// GetCate 查询分类列表
func GetProxies() ([]Proxies, error) {
	var proxies []Proxies
	err = db.Find(&proxies).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return []Proxies{}, err
	}
	return proxies, nil
}

func GetNamesFromAddrs(addrs []Proxies) ([]string, error) {
	var names []string
	for _, addr := range addrs {
		// 查询对应 server 的 name
		var name string
		if err := db.Model(&Proxies{}).Where("server = ? and inner_name = ? limit 1", addr.Server, addr.InnerName).Pluck("name", &name).Error; err != nil {
			fmt.Println("Failed to query names:", err)
			return []string{}, nil
		}
		names = append(names, name)
	}

	return names, nil
}
func GetConfsFromAddrs(addrs []string) ([]string, error) {
	var confs []string

	// 查询对应 server 的 name
	if err := db.Model(&Proxies{}).Where("server IN ?", addrs).Pluck("conf", &confs).Error; err != nil {
		fmt.Println("Failed to query names:", err)
		return []string{}, nil
	}
	return confs, nil
}

func GetServersFromType(typeName string) ([]Proxies, error) {
	var servers []Proxies
	// 查询对应 server 的 name
	if err := db.Where("proxie_type like ?", "%"+typeName+"%").Find(&servers).Error; err != nil {
		fmt.Println("Failed to query names:", err)
		return []Proxies{}, nil
	}
	return servers, nil
}
func GetServersFromServerAndInnerName(server, inner string) (Proxies, error) {

	var proxie Proxies
	if err := db.Where("server = ? and inner_name = ? limit 1", server, inner).Find(&proxie).Error; err != nil {
		fmt.Println("Failed to query names:", err)
		return Proxies{}, nil
	}
	return proxie, nil
}
