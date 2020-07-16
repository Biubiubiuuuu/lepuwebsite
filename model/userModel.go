package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
	"github.com/google/uuid"
)

// User 用户表
type User struct {
	Model
	Username       string    `gorm:"not null;unique;size:30;" json:"username"`  // 登录用户名（默认为注册手机号码）
	Telephone      string    `gorm:"not null;unique;size:30;" json:"telephone"` // 手机号码
	Password       string    `gorm:"size:50;" json:"-"`                         // 登录密码 （6-15位字符）
	Nickname       string    `gorm:"size:50;" json:"nickname"`                  // 姓名
	Sex            string    `gorm:"size:1;" json:"sex"`                        // 性别 0:未知 ｜ 1:男 ｜ 2:女 （空或其他默认未知）
	Landlinenumber string    `gorm:"size:30;" json:"landlinenumber"`            // 座机号码
	QQ             string    `gorm:"size:30;" json:"QQ"`                        // QQ
	Email          string    `gorm:"size:30;" json:"email"`                     // 邮箱
	IP             string    `gorm:"size:30;" json:"ip"`                        // 登录IP
	Token          string    `json:"token"`                                     // 授权令牌
	UUID           uuid.UUID `gorm:"size:50;" json:"uuid"`                      // 客户标识
	Type           string    `gorm:"size:1;" json:"-"`                          // 账号类型 0:用户 ｜ 1:管理员
	IsEnable       bool      `json:"is_enable"`                                 // 是否禁用 true | false （默认false，用户发布不良信息可禁用账号，并且所有信息对外不可见）
}

var db = mysql.GetMysqlDB()

// 用户注册
func (u *User) Register() error {
	return db.Create(&u).Error
}

// 修改用户信息
//  param id
func (u *User) Edit(args map[string]interface{}) error {
	return db.Model(u).Updates(args).Error
}

// 查询用户信息 by token
//  param token
//  return user,error
func (u *User) QueryByToken() error {
	return db.Where("token = ? AND ISNULL(token)=0 AND LENGTH(trim(token))>0", u.Token).First(&u).Error
}

// 查询用户信息 by telephone or username
//  param telephone or username
func (u *User) QueryByUsernameOrPhone() error {
	return db.Where("telephone = ? OR username = ?", u.Telephone, u.Username).First(&u).Error
}
