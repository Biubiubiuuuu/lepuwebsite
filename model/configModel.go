package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 网站信息配置
type SystemConfig struct {
	Model
	Favicon      string `gorm:"size:60;" json:"favicon"`                                          // 网站logo
	Hotline      string `gorm:"size:20;" json:"hotline"`                                          // 服务热线
	CusSerPhone  string `json:"cus_ser_phone"`                                                    // 客服手机号码
	Street       string `json:"street"`                                                           // 公司地址
	Advertphone  string `gorm:"size:20;" json:"advertphone"`                                      // 广告联系电话
	Copyright    string `json:"copyright"`                                                        // 版权所有
	CopyrightUrl string `gorm:"size:200;" json:"copyright_url"`                                   // 版权所有URL
	CaseNumber   string `json:"case_number"`                                                      // 备案号
	Links        []Link `gorm:"foreignkey:SystemConfigID;association_foreignkey:ID" json:"links"` // 友情链接
	IsDefault    bool   `json:"is_default"`                                                       // 是否设置默认
}

// 友情链接
type Link struct {
	ID             int64  `json:"-"`                    // ID
	Name           string `gorm:"size:50;" json:"name"` // 链接网站名称
	Url            string `gorm:"size:200;" json:"url"` // 链接网站URL
	SystemConfigID int64  `gorm:"INDEX" json:"-"`       // 网站信息配置ID
}

// 添加配置信息
func (s *SystemConfig) AddSystemConfig() error {
	db := mysql.GetMysqlDB()
	return db.Create(&s).Error
}

// 修改配置信息
func (s *SystemConfig) EditSystemConfig(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&s).Updates(args).Error
}

// 配置详情 by id
func (s *SystemConfig) QuerySystemConfigByID() error {
	db := mysql.GetMysqlDB()
	query := db.Table("system_config").Preload("Links").Select("system_config.*")
	return query.First(&s).Error
}

// 获取默认配置 by is_default
func (s *SystemConfig) QuerySystemConfigByDefault() error {
	db := mysql.GetMysqlDB()
	query := db.Table("system_config").Preload("Links").Select("system_config.*")
	return query.Where("is_default = true").First(&s).Error
}

// 配置列表
func QuerySystemConfig(pageSize int, page int) (count int, systemConfigs []SystemConfig) {
	db := mysql.GetMysqlDB()
	query := db.Table("system_config").Preload("Links").Select("system_config.*")
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&systemConfigs)
	return
}

// 设置默认
func (s *SystemConfig) EditDefaultSystemConfig() error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Model(&s).Update("is_default = true").Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(SystemConfig{}).Where("id IS NOT ?", s.ID).Update("is_default = false").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
