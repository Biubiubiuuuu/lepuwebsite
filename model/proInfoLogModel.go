package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

//  店铺转让跟单记录
type ProInfoLog struct {
	Model
	Content   string `json:"content"`                   // 跟单内容
	WithID    int64  `json:"with_id"`                   // 跟单人ID
	WithName  string `gorm:"size:50;" json:"with_name"` // 跟单人姓名
	ProInfoID int64  `json:"pro_info_id"`               // 物业信息ID
}

// 添加跟单记录
func (s *ProInfoLog) AddProInfoLog() error {
	db := mysql.GetMysqlDB()
	return db.Create(&s).Error
}

// 查看跟单记录
func QueryByProInfoID(id int64) (logs []ProInfoLog) {
	db := mysql.GetMysqlDB()
	db.Where("pro_info_id = ?", id).Find(&logs)
	return
}
