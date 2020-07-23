package model

import (
	"strconv"
	"strings"

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
	return db.First(&s).Error
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
func QueryStoreType(pageSize int, page int, name string, enable string) (count int, storeTypes []StoreType) {
	db := mysql.GetMysqlDB()
	query := db.Table("store_type").Select("store_type.*")
	if name != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(name)
		buf.WriteString("%")
		query = query.Where("name = ?", buf.String())
	}
	if enable != "" {
		boo, _ := strconv.ParseBool(enable)
		query = query.Where("enable = ?", boo)
	}
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&storeTypes)
	return
}

// 删除店铺类型，返回受影响行数
func DelStoreType(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&StoreType{}).RowsAffected
}
