package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  RentType 租金分类
type RentType struct {
	Model
	MinRent float64 `json:"min_rent"` // 最低租金（单位：元/月）
	MaxRent float64 `json:"max_rent"` // 最高租金（单位：元/月）
}

// 添加租金分类
func (r *RentType) AddRentType() error {
	db := mysql.GetMysqlDB()
	return db.Create(&r).Error
}

// 修改租金分类
func (r *RentType) EditRentType(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&r).Updates(args).Error
}

// 查询租金所在分组
func (r *RentType) QueryRentTypeByRent(rent float64) error {
	db := mysql.GetMysqlDB()
	return db.Where("min_rent < ? AND max_rent >= ?", rent, rent).First(&r).Error
}

// 查询租金分类
func (r *RentType) QueryRentType() (rentTypes []RentType) {
	db := mysql.GetMysqlDB()
	db.Find(&rentTypes)
	return
}

// 删除租金分类(可批量)
func (r *RentType) DeleteRentType(ids []int64) error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Unscoped().Delete("id in (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
