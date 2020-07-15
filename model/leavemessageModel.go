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
func (l *LeaveMessage) QueryLeaveMessage(pageSize int, page int) (leaveMessages []LeaveMessage) {
	db := mysql.GetMysqlDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&leaveMessages)
	return
}

// 查看留言总记录数
func (l *LeaveMessage) QueryLeaveMessageCount() (count int) {
	db := mysql.GetMysqlDB()
	db.Model(&LeaveMessage{}).Count(&count)
	return
}
