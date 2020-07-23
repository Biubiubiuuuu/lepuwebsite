package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 岗位
type Post struct {
	Model
	Name   string `gorm:"size:50;unique" json:"name"` // 岗位名称
	Code   string `gorm:"size:50;unique" json:"code"` // 岗位编码
	Sort   int64  `json:"sort"`                       // 显示排序
	Enable bool   `json:"enable"`                     // 是否启用
}

// 添加岗位
func (p *Post) AddPost() error {
	db := mysql.GetMysqlDB()
	return db.Create(&p).Error
}

// 修改岗位
func (p *Post) EditPost(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&p).Updates(args).Error
}

// 查询岗位信息 by id
func (p *Post) QueryPostByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&p, p.ID).Error
}

// 查询岗位信息 by name
func (p *Post) QueryPostByName() error {
	db := mysql.GetMysqlDB()
	return db.Where("name = ?", p.Name).First(&p).Error
}

// 查询岗位信息 by code
func (p *Post) QueryPostByCode() error {
	db := mysql.GetMysqlDB()
	return db.Where("code = ?", p.Code).First(&p).Error
}

// 查询所有岗位
func QueryPosts(pageSize int, page int, args map[string]interface{}) (count int, posts []Post) {
	db := mysql.GetMysqlDB()
	query := db.Table("post").Select("post.*")
	if v, ok := args["name"]; ok && v.(string) != "" {
		query = query.Where("name = %?%", v.(string))
	}
	if v, ok := args["enable"]; ok {
		query = query.Where("enable = ?", v.(string))
	}
	if v, ok := args["code"]; ok && v.(string) != "" {
		query = query.Where("code = %?%", v.(string))
	}
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&posts)
	return
}

// 删除岗位，返回受影响行数
func DelPosts(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Post{}).RowsAffected
}
