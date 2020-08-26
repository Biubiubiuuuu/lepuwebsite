package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// LeaveMessage 留言
type LeaveMessage struct {
	Model
	Content   string `gorm:"size:200;" json:"content"`  // 留言内容
	Address   string `gorm:"size:100;" json:"address"`  // 详细地址
	Telephone string `gorm:"size:20;" json:"telephone"` // 联系手机
	Nickname  string `gorm:"size:20;" json:"nickname"`  // 联系人
}

// 添加留言
func (l *LeaveMessage) AddLeaveMessage() error {
	db := mysql.GetMysqlDB()
	return db.Create(&l).Error
}

// 查看留言
func QueryLeaveMessage(pageSize int, page int) (count int, leaveMessages []LeaveMessage) {
	db := mysql.GetMysqlDB()
	query := db.Table("leave_message").Select("leave_message.*")
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&leaveMessages)
	return
}

// 查看留言详情
func (l *LeaveMessage) QueryLeaveMessageByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&l, l.ID).Error
}
