package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 广告
type Advert struct {
	Model
	Sort           int64     `json:"sort"`                // 显示排序
	StartTime      time.Time `json:"start_time"`          // 展示开始时间
	EndTime        time.Time `json:"end_time"`            // 展示结束时间
	Hot            bool      `json:"hot"`                 // 首页最热推广
	Floor          bool      `json:"floor"`               // F楼
	Type           string    `gorm:"size:1;" json:"type"` // 信息列表推广 1-一栏四分之一图片广告 | 2-二栏四分之一图片广告 | 3-三栏重点推荐 | 4-五栏框架广告
	PropertyInfoID int64     `json:"property_info_id"`    // 物业ID
}

// 添加广告
func (a *Advert) AddAdvert() error {
	db := mysql.GetMysqlDB()
	return db.Create(&a).Error
}

// 查询广告
func QueryAdvert(pageSize int, page int, args map[string]interface{}) (count int, adverts []Advert) {
	db := mysql.GetMysqlDB()
	query := db.Table("advert").Select("advert.*")
	if v, ok := args["type"]; ok && v.(string) != "" {
		query = query.Where("type = ?", v.(string))
	}
	if v, ok := args["hot"]; ok && v.(string) != "" {
		hot, _ := strconv.ParseBool(v.(string))
		query = query.Where("hot = ?", hot)
	}
	if v, ok := args["floor"]; ok && v.(string) != "" {
		floor, _ := strconv.ParseBool(v.(string))
		query = query.Where("floor = ?", floor)
	}
	fmt.Println(time.Now())
	query = query.Where("TO_DAYS(end_time) >= TO_DAYS(?)", time.Now())
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&adverts)
	return
}

// 查询广告详情
func (a *Advert) QueryAdvertByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&a).Error
}

// 修改广告
func (a *Advert) EditAdvertByID(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&a).Updates(args).Error
}

// 删除广告，返回受影响行数
func DelAdvert(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Advert{}).RowsAffected
}
