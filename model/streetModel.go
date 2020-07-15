package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  street 街道
type Street struct {
	Model
	Name         string `gorm:"not null;unique;size:10;" json:"name"`   // 街道名称
	Code         string `gorm:"not null;unique;size:10;" json:"code"`   // 街道代码
	DistrictCode string `gorm:"not null;size:10;" json:"district_code"` // 区代码
}

// 根据街道代码找街道
func (s *Street) QueryStreetByCode() error {
	db := mysql.GetMysqlDB()
	return db.Where("code = ?", s.Code).First(&s).Error
}

// 根据区查找街道
func (s *Street) QueryStreetByDistrictCode() (streets []Street) {
	db := mysql.GetMysqlDB()
	db.Where("district_code = ?", s.DistrictCode).Find(&streets)
	return
}