package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  District 区
type District struct {
	Model
	Name     string `gorm:"not null;size:10;" json:"name"`        // 区名称
	Code     string `gorm:"not null;unique;size:10;" json:"code"` // 区代码
	CityCode string `gorm:"not null;size:10;" json:"city_code"`   // 城市代码
}

// 根据区代码查找区
func (d *District) QueryDistrictByCode() error {
	db := mysql.GetMysqlDB()
	return db.Where("code = ?", d.Code).First(&d).Error
}

// 根据城市代码查找区
func (d *District) QueryDistrictByCityCode() (districts []District) {
	db := mysql.GetMysqlDB()
	db.Where("city_code = ?", d.CityCode).Find(&districts)
	return
}

//  查找所有区
func (d *District) QueryDistrict() (districts []District) {
	db := mysql.GetMysqlDB()
	db.Find(&districts)
	return
}

// 根据区代码和城市代码查找区
func (d *District) QueryDistrictByCodeAndCity() error {
	db := mysql.GetMysqlDB()
	return db.Where("code = ? AND city_code = ?", d.Code, d.CityCode).First(&d).Error
}
