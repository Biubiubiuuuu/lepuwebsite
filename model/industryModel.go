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

// 删除行业(可批量)
func (i *Industry) DeleteIndustry(ids []int64) error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Unscoped().Delete("id in (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 查询已启用适合经营范围
func (i *Industry) QueryEnableIndustryRange() (industrys []Industry) {
	db := mysql.GetMysqlDB()
	db.Where("is_enable = ? AND parent_id IN (SELECT id FROM industry WHERE parent_id = 0)", true).Order("sort desc").Find(&industrys)
	return
}
