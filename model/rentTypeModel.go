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

// 查询租金分类 by id
func (r *RentType) QueryRentTypeInfoById() error {
	db := mysql.GetMysqlDB()
	return db.First(&r).Error
}

// 删除面积分类，返回受影响行数
func DelRentType(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&RentType{}).RowsAffected
}

// 查询已添加面积范围最大值
func (r *RentType) QueryMaxRent() error {
	db := mysql.GetMysqlDB()
	return db.Raw("SELECT * FROM rent_type WHERE max_rent = (SELECT MAX(max_rent) FROM rent_type)").Scan(&r).Error
}
