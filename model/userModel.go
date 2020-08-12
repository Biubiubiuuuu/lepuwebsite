package model

import (
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
	"github.com/google/uuid"
)

// User 用户表
type User struct {
	Model
	Username       string    `gorm:"not null;unique;size:30;" json:"username"`                               // 登录用户名（默认为注册手机号码）
	Telephone      string    `gorm:"not null;unique;size:30;" json:"telephone"`                              // 手机号码
	Password       string    `gorm:"size:50;" json:"-"`                                                      // 登录密码 （6-15位字符）
	Nickname       string    `gorm:"size:50;" json:"nickname"`                                               // 姓名
	Sex            string    `gorm:"size:1;" json:"sex"`                                                     // 性别 0:未知 ｜ 1:男 ｜ 2:女 （空或其他默认未知）
	Landlinenumber string    `gorm:"size:30;" json:"landlinenumber"`                                         // 座机号码
	QQ             string    `gorm:"size:30;" json:"QQ"`                                                     // QQ
	Email          string    `gorm:"size:30;" json:"email"`                                                  // 邮箱
	IP             string    `gorm:"size:30;" json:"ip"`                                                     // 登录IP
	Token          string    `json:"token"`                                                                  // 授权令牌
	UUID           uuid.UUID `gorm:"size:50;" json:"uuid"`                                                   // 客户标识
	Type           string    `gorm:"size:1;" json:"-"`                                                       // 账号类型 0:用户 ｜ 1:管理员
	UserInfo       UserInfo  `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"user_info,omitempty"` // 账号类型为管理员时，绑定其岗位、角色等信息
	IsEnable       bool      `json:"is_enable"`                                                              // 是否禁用（默认false，用户发布不良信息可禁用账号，并且所有信息对外不可见）
}

// 用户信息表
type UserInfo struct {
	ID             int64  `json:"-"`                                         // ID
	UserID         int64  `gorm:"INDEX" json:"-"`                            // 用户ID
	DepartmentID   int64  `json:"department_id,omitempty"`                   // 部门ID
	DepartmentName string `gorm:"size:50;" json:"department_name,omitempty"` // 部门名称
	PostID         int64  `json:"post_id,omitempty"`                         // 岗位ID
	PostName       string `gorm:"size:50;" json:"post_name,omitempty"`       // 岗位名称
	RoleID         int64  `json:"role_id"`                                   // 角色ID
	RoleName       string `gorm:"size:50;" json:"role_name"`                 // 角色名称
}

// 用户注册
func (u *User) Register() error {
	db := mysql.GetMysqlDB()
	return db.Create(&u).Error
}

// 修改用户信息
//  param id
func (u *User) Edit(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	db.Model(&User{}).Association("UserInfo").Replace(u.UserInfo)
	return db.Model(&u).Updates(args).Error
}

// 查询用户信息 by token
//  param token
//  return user,error
func (u *User) QueryByToken() error {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	return query.Where("token = ? AND ISNULL(token)=0 AND LENGTH(trim(token))>0", u.Token).First(&u).Error
}

// 查询用户信息 by telephone or username
//  param telephone or username
func (u *User) QueryByUsernameOrPhone() error {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	return query.Where("telephone = ? OR username = ?", u.Telephone, u.Username).First(&u).Error
}

// 查询用户信息 by telephone
//  param telephone
func (u *User) QueryByPhone() error {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	return query.Where("telephone = ?", u.Telephone).First(&u).Error
}

// 查询员工信息 by id
//  param telephone
func (u *User) QueryEmployeeById() error {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	return query.Where("type = 1").First(&u).Error
}

// 查询用户信息 by id
//  param telephone
func (u *User) QueryUserByID() error {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	return query.First(&u).Error
}

// 查询用户信息 by  username
//  param username
func (u *User) QueryByUsername() error {
	db := mysql.GetMysqlDB()
	return db.Where("username = ?", u.Username).First(&u).Error
}

// 查询用户信息是否关联部门
func QueryUserByDepartmentID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("user_info")
	query = query.Joins("left user_info on user_info.user_id = user.id")
	var users []User
	if count := query.Where("user_info.department_id in (?)", ids).Find(&users).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询用户信息是否关联岗位
func QueryUserByPostID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("user_info")
	query = query.Joins("left user_info on user_info.user_id = user.id")
	var users []User
	if count := query.Where("user_info.post_id in (?)", ids).Find(&users).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询用户信息是否关联角色
func QueryUserByRoleID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("user_info")
	query = query.Joins("left user_info on user_info.user_id = user.id")
	var users []User
	if count := query.Where("user_info.role_id in (?)", ids).Find(&users).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 删除员工，返回受影响行数
func DelEmployee(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&User{}).RowsAffected
}

// 查询员工信息
func QueryUser(pageSize int, page int, args map[string]interface{}) (count int, users []User) {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("UserInfo")
	if v, ok := args["username"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("user.username like ?", buf.String())
	}
	if v, ok := args["nickname"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("user.nickname like ?", buf.String())
	}
	if v, ok := args["enable"]; ok && v.(string) != "" {
		enable, _ := strconv.ParseBool(v.(string))
		query = query.Where("user.is_enable = ?", enable)
	}
	if v, ok := args["telephone"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("user.telephone like ?", buf.String())
	}
	query = query.Where("type = 1")
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
	return
}
