package model

import (
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 部门
type Department struct {
	Model
	Name     string `gorm:"size:50;unique" json:"name"` // 部门名称
	Sort     int64  `json:"sort"`                       // 显示排序
	Leading  string `gorm:"size:50;" json:"leading"`    // 负责人
	Phone    string `gorm:"size:20;" json:"phone"`      // 联系电话
	Email    string `gorm:"size:30;" json:"email"`      // 邮箱
	Enable   bool   `json:"enable"`                     // 是否启用
	ParentID int64  `gorm:"default:0" json:"parent_id"` // 上级ID 0为最顶级
}

// 添加部门
func (d *Department) AddDepartment() error {
	db := mysql.GetMysqlDB()
	return db.Create(&d).Error
}

// 修改部门 by ID
func (d *Department) EditDepartmentByID(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&d).Updates(args).Error
}

// 查看部门详情 by id
func (d *Department) QueryDepartmentByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&d).Error
}

// 查询上级部门下的部门详情 by name and parent_id
func (d *Department) QueryDepartmentByNameAndParentID() error {
	db := mysql.GetMysqlDB()
	return db.Where("name = ? AND parent_id = ?", d.Name, d.ParentID).First(&d).Error
}

// 查询已启用部门
func QueryEnableDepartments() (departments []Department) {
	db := mysql.GetMysqlDB()
	db.Where("enable = ?", true).Find(&departments)
	return
}

// 查询所有部门
func QueryDepartments(pageSize int, page int, name string, enable string) (count int, departments []Department) {
	db := mysql.GetMysqlDB()
	query := db.Table("department").Select("department.*")
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
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&departments)
	return
}

// 删除部门，返回受影响行数
func DelDepartments(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Department{}).RowsAffected
}
