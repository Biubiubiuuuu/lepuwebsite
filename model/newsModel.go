package model

// 行业资讯
type News struct {
	Model
	Title   string `gorm:"size:150;unique" json:"title"` // 标题
	Source  string `gorm:"size:20" json:"source"`        // 来源
	Content string `gorm:"size:5000" json:"content"`     // 内容
}
