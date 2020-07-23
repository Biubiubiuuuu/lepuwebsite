package model

import "github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"

// 菜单
type Menu struct {
	Model
	Title      string `gorm:"size:30;" json:"title"`      // 菜单标题
	Sort       int64  `json:"sort"`                       // 显示排序
	Icon       string `gorm:"size:50" json:"icon"`        // 菜单图标
	RouterName string `gorm:"size:50" json:"router_name"` // 路由名称
	RouterUrl  string `gorm:"size:100" json:"router_url"` // 路由地址
	Enable     bool   `json:"enable"`                     // 是否启用
	ParentID   int64  `gorm:"default:0" json:"parent_id"` // 上级ID 0为最顶级
}

// 创建菜单
func (m *Menu) AddMenu() error {
	db := mysql.GetMysqlDB()
	return db.Create(&m).Error
}

// 修改菜单
func (m *Menu) EditMenu(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&m).Updates(args).Error
}

// 获取所有菜单
func QueryMenus() (menus []Menu) {
	db := mysql.GetMysqlDB()
	db.Find(&menus)
	return
}

// 获取菜单详情
func (m *Menu) QueryMenuByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&m).Error
}

// 删除菜单，返回受影响行数
func DelMebus(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Menu{}).RowsAffected
}
