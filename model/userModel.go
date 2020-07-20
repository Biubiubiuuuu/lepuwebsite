package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
	"github.com/google/uuid"
)

// User 用户表
type User struct {
	Model
	Username       string    `gorm:"not null;unique;size:30;" json:"username"`                     // 登录用户名（默认为注册手机号码）
	Telephone      string    `gorm:"not null;unique;size:30;" json:"telephone"`                    // 手机号码
	Password       string    `gorm:"size:50;" json:"-"`                                            // 登录密码 （6-15位字符）
	Nickname       string    `gorm:"size:50;" json:"nickname"`                                     // 姓名
	Sex            string    `gorm:"size:1;" json:"sex"`                                           // 性别 0:未知 ｜ 1:男 ｜ 2:女 （空或其他默认未知）
	Landlinenumber string    `gorm:"size:30;" json:"landlinenumber"`                               // 座机号码
	QQ             string    `gorm:"size:30;" json:"QQ"`                                           // QQ
	Email          string    `gorm:"size:30;" json:"email"`                                        // 邮箱
	IP             string    `gorm:"size:30;" json:"ip"`                                           // 登录IP
	Token          string    `json:"token"`                                                        // 授权令牌
	UUID           uuid.UUID `gorm:"size:50;" json:"uuid"`                                         // 客户标识
	Type           string    `gorm:"size:1;" json:"-"`                                             // 账号类型 0:用户 ｜ 1:管理员
	UserInfo       UserInfo  `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"user_info"` // 账号类型为管理员时，绑定其岗位、角色等信息
	IsEnable       bool      `json:"is_enable"`                                                    // 是否禁用（默认false，用户发布不良信息可禁用账号，并且所有信息对外不可见）
}

// 用户信息表
type UserInfo struct {
	ID             int64      `json:"-"`                                                                 // ID
	UserID         int64      `gorm:"INDEX" json:"-"`                                                    // 用户ID
	UserRoles      []UserRole `gorm:"foreignkey:UserInfoID;association_foreignkey:ID" json:"user_roles"` // 用户角色
	DepartmentID   int64      `json:"department_id"`                                                     // 部门ID
	DepartmentName string     `gorm:"size:50;" json:"department_name"`                                   // 部门名称
	PostID         int64      `json:"post_id"`                                                           // 岗位ID
	PostName       string     `gorm:"size:50;" json:"post_name"`                                         // 岗位名称
}

// 用户角色
type UserRole struct {
	ID         int64  `json:"-"`                         // ID
	RoleID     int64  `json:"role_id"`                   // 角色ID
	RoleName   string `gorm:"size:50;" json:"role_name"` // 角色名称
	UserInfoID int64  `gorm:"INDEX" json:"-"`            // ID
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
	return db.Model(u).Updates(args).Error
}

// 查询用户信息 by token
//  param token
//  return user,error
func (u *User) QueryByToken() error {
	db := mysql.GetMysqlDB()
	return db.Where("token = ? AND ISNULL(token)=0 AND LENGTH(trim(token))>0", u.Token).First(&u).Error
}

// 查询用户信息 by telephone or username
//  param telephone or username
func (u *User) QueryByUsernameOrPhone() error {
	db := mysql.GetMysqlDB()
	return db.Where("telephone = ? OR username = ?", u.Telephone, u.Username).First(&u).Error
}

// 查询用户信息是否关联部门
func QueryUserByDepartmentID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	query := db.Table("user").Preload("user_info").Preload("user_roles")
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
	query := db.Table("user").Preload("user_info").Preload("user_roles")
	query = query.Joins("left user_info on user_info.user_id = user.id")
	var users []User
	if count := query.Where("user_info.post_id in (?)", ids).Find(&users).RowsAffected; count > 0 {
		return true
	}
	return false
}
