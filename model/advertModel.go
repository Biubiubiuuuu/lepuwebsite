package model

// 广告
type Advert struct {
	Model
	Sort           int64  `json:"sort"`                // 显示排序
	StartTime      string `json:"start_time"`          // 展示开始时间
	EndTime        string `json:"end_time"`            // 展示结束时间
	Type           string `gorm:"size:1;" json:"type"` // 广告类型 0-F楼广告 1-轮播广告 2-推荐广告
	PropertyInfoID int64  `json:"property_info_id"`    // 物业ID
}
