package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 举报信息
type Report struct {
	Model
	UserID       int64  `json:"user_id"`                  // 举报者ID
	Nickname     string `gorm:"size:50;" json:"nickname"` // 举报者
	ProInfoID    int64  `json:"property_info_id"`         // 物业信息ID
	ProInfoTitle string `gorm:"size:255;" json:"title"`   // 标题
	Content      string `gorm:"size:255;" json:"content"` // 举报内容
}

// 添加举报信息
func (r *Report) AddReport() error {
	db := mysql.GetMysqlDB()
	return db.Create(&r).Error
}

// 查看举报信息
func QueryReport(pageSize int, page int) (count int, reports []Report) {
	db := mysql.GetMysqlDB()
	db.Table("report").Count(&count)
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&reports)
	return
}

// 查看举报信息详情
func (r *Report) QueryReportByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&r, r.ID).Error
}
