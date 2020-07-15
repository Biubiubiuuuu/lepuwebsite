package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// ShopTransferLog 店铺转让跟单记录
type ShopTransferLog struct {
	Model
	Content        string `json:"content"`                   // 跟单内容
	WithID         int64  `json:"with_id"`                   // 跟单人ID
	WithName       string `gorm:"size:50;" json:"with_name"` // 跟单人姓名
	ShopTransferID int64  `json:"shop_transfer_id"`          // 店铺转让ID
}

// 添加跟单记录
func (s *ShopTransferLog) AddShopTransferLog() error {
	db := mysql.GetMysqlDB()
	return db.Create(&s).Error
}

// 查看跟单记录
func (s *ShopTransferLog) QueryByShopTransferID() (logs []ShopTransferLog) {
	db := mysql.GetMysqlDB()
	db.Where("shop_transfer_id = ?", s.ShopTransferID).Find(&logs)
	return
}
