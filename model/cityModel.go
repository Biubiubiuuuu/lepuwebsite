package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  City 城市
type City struct {
	Model
	Name         string `gorm:"not null;unique;size:10;" json:"name"`   // 城市名称
	Code         string `gorm:"not null;unique;size:10;" json:"code"`   // 城市代码
	ProvinceCode string `gorm:"not null;size:10;" json:"province_code"` // 省代码
}

// 根据城市代码查询城市
func (c *City) QueryCitysByCode() error {
	db := mysql.GetMysqlDB()
	return db.Where("code = ?", c.Code).First(&c).Error
}

// 根据省代码查询城市
func (c *City) QueryCitysByProvinceCode() (citys []City) {
	db := mysql.GetMysqlDB()
	db.Where("province_code = ?", c.ProvinceCode).Find(&citys)
	return
}
