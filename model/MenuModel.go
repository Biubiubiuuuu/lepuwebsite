package model

// 菜单
type Menu struct {
	Model
	Title      string `gorm:"size:30;unique" json:"title"` // 菜单标题
	Sort       int64  `json:"sort"`                        // 显示排序
	Type       string `gorm:"size:1" json:"type"`          // 菜单类型 0-目录 | 1-菜单
	Icon       string `gorm:"size:50" json:"icon"`         // 菜单图标
	RouterName string `gorm:"size:50" json:"router_name"`  // 路由名称
	RouterUrl  string `gorm:"size:100" json:"router_url"`  // 路由地址
	Enable     bool   `json:"enable"`                      // 是否启用
	ParentID   int64  `gorm:"default:0" json:"parent_id"`  // 上级ID 0为最顶级
}
