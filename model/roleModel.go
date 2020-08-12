package model

import (
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 角色
type Role struct {
	Model
	Name       string      `gorm:"size:50;unique" json:"name"`                                     // 角色名称
	Sort       int64       `json:"sort"`                                                           // 显示排序
	Enable     bool        `json:"enable"`                                                         // 是否启用
	MenuPowers []MenuPower `gorm:"foreignkey:RoleID;association_foreignkey:ID" json:"menu_powers"` // 菜单权限
}

// 角色菜单权限
type MenuPower struct {
	ID        int64  `json:"id"`                         // ID
	MenuID    int64  `json:"menu_id"`                    // 菜单ID
	MenuTitle string `gorm:"size:30;" json:"menu_title"` // 菜单标题
	RoleID    int64  `json:"-"`                          // 角色ID
}

// 创建角色
func (r *Role) AddRole() error {
	db := mysql.GetMysqlDB()
	return db.Create(&r).Error
}

// 修改角色 by id
//func (r *Role) Editrole(args map[string]interface{}) error {
//	db := mysql.GetMysqlDB()
//	return db.Model(&r).Updates(args).Error
//}

// 查询角色详情 by id
func (r *Role) QueryRoleByID() error {
	db := mysql.GetMysqlDB()
	query := db.Table("role").Preload("MenuPowers")
	return query.First(&r).Error
}

// 查询角色详情 by name
func (r *Role) QueryRoleByName() error {
	db := mysql.GetMysqlDB()
	query := db.Table("role").Preload("MenuPowers")
	return query.Where("name = ?", r.Name).First(&r).Error
}

// 删除角色，返回受影响行数
func DelRoles(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Role{}).RowsAffected
}

// 修改角色信息
func (r *Role) EditRole(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	db.Model(&r).Association("MenuPowers").Replace(r.MenuPowers)
	return db.Model(&r).Updates(args).Error
}

// 查询角色是否关联菜单
func QueryRoleByMenuID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	var menu_powers []MenuPower
	if count := db.Where("role_id in (?)", ids).Find(&menu_powers).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询角色
func QueryRole(pageSize int, page int, name string, enable string) (count int, roles []Role) {
	db := mysql.GetMysqlDB()
	query := db.Table("role").Preload("MenuPowers")
	if name != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(name)
		buf.WriteString("%")
		query = query.Where("name like ?", buf.String())
	}
	if enable != "" {
		boo, _ := strconv.ParseBool(enable)
		query = query.Where("enable = ?", boo)
	}
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&roles)
	return
}

// 查询用户所拥有的角色菜单
func QueryMenuByUser(id int64) (count int, menus []Menu) {
	db := mysql.GetMysqlDB()
	query := db.Table("menu").Select("menu.*")
	query = query.Where("menu.id in (SELECT menu_id FROM menu_power WHERE role_id = ?)", id)
	query.Count(&count)
	query.Find(&menus)
	return
}
