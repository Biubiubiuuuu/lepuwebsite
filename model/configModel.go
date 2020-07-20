package model

// 网站信息配置
type SystemConfig struct {
	Model
	Favicon      string `gorm:"size:60;" json:"favicon"`                                  // 网站logo
	Hotline      string `gorm:"size:20;" json:"hotline"`                                  // 服务热线
	Street       string `json:"street"`                                                   // 公司地址
	Advertphone  string `gorm:"size:20;" json:"advertphone"`                              // 广告联系电话
	Copyright    string `json:"copyright"`                                                // 版权所有
	CopyrightUrl string `gorm:"size:200;" json:"copyright_url"`                           // 版权所有URL
	Links        []Link `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"links"` // 友情链接
}

// 友情链接
type Link struct {
	ID              int64  `json:"-"`                    // ID
	Name            string `gorm:"size:50;" json:"name"` // 链接网站名称
	Url             string `gorm:"size:200;" json:"url"` // 链接网站URL
	WebsiteConfigID int64  `gorm:"INDEX" json:"-"`       // 网站信息配置ID
}
