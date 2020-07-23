package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  Province 省
type Province struct {
	Model
	Name string `gorm:"not null;unique;size:10;" json:"name"` // 省名称
	Code string `gorm:"not null;unique;size:10;" json:"code"` // 省代码
}

// 根据省代码查询省
func (p *Province) QueryProvinceByCode() error {
	db := mysql.GetMysqlDB()
	return db.Where("code =?", p.Code).First(&p).Error
}

// 查询省
func (p *Province) QueryProvinces() (provinces []Province) {
	db := mysql.GetMysqlDB()
	db.Find(&provinces)
	return
}
