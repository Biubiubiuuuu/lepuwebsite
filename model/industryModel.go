package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  Industry 行业
type Industry struct {
	Model
	Name     string `gorm:"not null;unique;size:20;" json:"name"` // 行业名称
	Sort     int64  `json:"sort"`                                 // 行业排序 越大越靠前
	IsEnable bool   `json:"is_enable"`                            // 是否启用 true | false
	ParentID int64  `gorm:"not null;default:0;" json:"parent_id"` // 父类型ID
}

// 添加行业
func (i *Industry) AddIndustry() error {
	db := mysql.GetMysqlDB()
	return db.Create(&i).Error
}

// 修改行业
func (i *Industry) EditIndustry(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&i).Updates(args).Error
}

// 查询行业 by id
func (i *Industry) QueryIndustryByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&i, i.ID).Error
}

// 查询行业 by name
func (i *Industry) QueryIndustryByName() error {
	db := mysql.GetMysqlDB()
	return db.Where("name =?", i.Name).First(&i).Error
}

// 查询已启用行业类型
func (i *Industry) QueryEnableIndustry() (industrys []Industry) {
	db := mysql.GetMysqlDB()
	db.Where("is_enable = ?", true).Order("sort desc").Find(&industrys)
	return
}

// 查询所有行业类型
func (i *Industry) QueryIndustry() (industrys []Industry) {
	db := mysql.GetMysqlDB()
	db.Order("sort desc").Find(&industrys)
	return
}

// 删除行业类型，返回受影响行数
func DelIndustry(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Industry{}).RowsAffected
}

// 查询已启用最上级行业类型
func QueryEnableIndustryByParentID() (industrys []Industry) {
	db := mysql.GetMysqlDB()
	db.Where("is_enable = true AND parent_id = 0").Order("sort desc").Find(&industrys)
	return
}
