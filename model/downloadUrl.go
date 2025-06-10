package model

type downloadUrl struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	DownloadUrl string `gorm:"type:varchar(128);not null" json:"download_url"`
}

func GetDownloadUrl() (string, error) {
	var url downloadUrl
	// 查询第一个记录
	if err := db.First(&url).Error; err != nil {
		return "", err // 返回错误，如果查询失败
	}
	return url.DownloadUrl, nil
}
