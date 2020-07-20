package model

import "github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"

// 角色
type Role struct {
	Model
	Name       string      `gorm:"size:50;unique" json:"name"`                                           // 角色名称
	Sort       int64       `json:"sort"`                                                                 // 显示排序
	Enable     bool        `json:"enable"`                                                               // 是否启用
	MenuPowers []MenuPower `gorm:"foreignkey:MenuParentID;association_foreignkey:ID" json:"menu_powers"` // 菜单权限
}

// 角色菜单权限
type MenuPower struct {
	ID        int64  `json:"id"`                         // ID
	MenuID    int64  `json:"menu_id"`                    // 菜单ID
	MenuTitle string `gorm:"size:30;" json:"menu_title"` // 菜单标题
}

// 创建角色
func (r *Role) AddRole() error {
	db := mysql.GetMysqlDB()
	return db.Create(&r).Error
}

// 修改角色 by id
func (r *Role) Editrole(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&r).Updates(args).Error
}

// 查询角色详情 by id
func (r *Role) QueryRoleByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&r, r.ID).Error
}

// 查询角色详情 by name
func (r *Role) QueryRoleByName() error {
	db := mysql.GetMysqlDB()
	return db.Where("name = ?", r.Name).First(&r).Error
}
