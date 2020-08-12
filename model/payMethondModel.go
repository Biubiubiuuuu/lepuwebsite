package model

import "github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"

// 付款方式
type PayMethond struct {
	Model
	Name string `json:"name"` // 付款方式名称
	Card string `json:"card"` // 付款卡号
}

// 添加付款方式
func (p *PayMethond) AddPayMethond() error {
	db := mysql.GetMysqlDB()
	return db.Create(&p).Error
}

// 修改付款方式
func (p *PayMethond) EditPayMethond(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&p).Updates(args).Error
}

// 查询所有付款方式
func QueryPayMethond() (payMethonds []PayMethond) {
	db := mysql.GetMysqlDB()
	db.Find(&payMethonds)
	return
}

// 付款方式详情
func (p *PayMethond) QueryPayMethondByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&p).Error
}

// 删除付款方式，返回受影响行数
func DelPayMethond(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&PayMethond{}).RowsAffected
}
