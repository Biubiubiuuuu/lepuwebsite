package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  StoreTypeType 店铺类型
type StoreType struct {
	Model
	Name     string `gorm:"not null;unique;size:20;" json:"name"` // 类型名称
	Sort     int64  `json:"sort"`                                 // 类型排序 越大越靠前
	IsEnable bool   `json:"is_enable"`                            // 是否启用 true | false
}

// 添加店铺类型
func (s *StoreType) AddStoreType() error {
	db := mysql.GetMysqlDB()
	return db.Create(&s).Error
}

// 修改店铺类型
func (s *StoreType) EditStoreType(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&s).Updates(args).Error
}

// 查询店铺 by id
func (s *StoreType) QueryStoreTypeByID() error {
	db := mysql.GetMysqlDB()
	return db.Model(&s).First(&s).Error
}

// 查询店铺 by name
func (s *StoreType) QueryStoreTypeByName() error {
	db := mysql.GetMysqlDB()
	return db.Where("name =?", s.Name).First(&s).Error
}

// 查询已启用店铺类型
func (s *StoreType) QueryEnableStoreType() (storeTypes []StoreType) {
	db := mysql.GetMysqlDB()
	db.Where("is_enable = ?", true).Order("sort desc").Find(&storeTypes)
	return
}

// 查询所有店铺类型
func (s *StoreType) QueryStoreType() (storeTypes []StoreType) {
	db := mysql.GetMysqlDB()
	db.Order("sort desc").Find(&storeTypes)
	return
}

// 删除店铺类型(可批量)
func (s *StoreType) DeleteStoreType(ids []int64) error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Unscoped().Delete("id in (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
