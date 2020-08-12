package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 付款
type PayInfo struct {
	Model
	Name             string    `json:"name"`              // 收款人
	PayeeID          int64     `json:"payee_id"`          // 业绩归属ID
	Payee            string    `json:"payee"`             // 业绩归属人名称
	PayMethondID     int64     `json:"pay_methond_id"`    // 付款方式ID
	PayMethond       string    `json:"pay_methond"`       // 付款方式名称
	PayTime          time.Time `json:"pay_time"`          // 收款时间
	PayStatus        string    `json:"pay_status"`        // 收款情况
	ActualAmount     float64   `json:"actual_amount"`     // 实收金额
	ReceivableAmount float64   `json:"receivable_amount"` // 应收金额
	Invoice          bool      `json:"invoice"`           // 发票
	Remake           string    `json:"remake"`            // 备注说明
	ProInfoID        int64     `json:"pro_info_id"`       // 物业ID
	ProInfoTitle     string    `json:"pro_info_title"`    // 物业标题
	ProInfoNickname  string    `json:"nickname"`          // 联系人
}

// 添加收款
func (p *PayInfo) AddPayInfo() error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	pro := PropertyInfo{}
	pro.ID = p.ProInfoID
	if err := pro.EditPropertyInfoByID(map[string]interface{}{"status": true}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 查询收款详情
func (p *PayInfo) QueryPayInfoByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&p).Error
}

// 查询收款详情
func (p *PayInfo) QueryPayInfoByProInfoID() error {
	db := mysql.GetMysqlDB()
	return db.Where("pro_info_id = ?", p.ProInfoID).First(&p).Error
}

// 收款列表
func QueryPayInfo(pageSize int, page int, args map[string]interface{}) (count int, payInfos []PayInfo) {
	db := mysql.GetMysqlDB()
	query := db.Table("pay_info").Select("pay_info.*")
	if v, ok := args["payee_id"]; ok && v.(string) != "" {
		payee_id, _ := strconv.ParseInt(v.(string), 10, 64)
		query = query.Where("payee_id = ?", payee_id)
	}
	if v, ok := args["pay_methond_id"]; ok && v.(string) != "" {
		pay_methond_id, _ := strconv.ParseInt(v.(string), 10, 64)
		query = query.Where("pay_methond_id = ?", pay_methond_id)
	}
	// 月份
	if v, ok := args["pay_month"]; ok && v.(string) != "" {
		query = query.Where("MONTH(pay_time) = ?", v.(string))
	}
	// 年
	if v, ok := args["pay_year"]; ok && v.(string) != "" {
		query = query.Where("YEAR(pay_time) = ?", v.(string))
	}
	if v, ok := args["name"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("name = ?", buf.String())
	}
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&payInfos)
	return
}

// 修改收款
func (p *PayInfo) EditPayInfo(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Model(&p).Updates(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	pro := PropertyInfo{}
	pro.ID = p.ProInfoID
	if err := pro.EditPropertyInfoByID(map[string]interface{}{"status": true}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 删除付款，返回受影响行数
func DelPayInfo(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&PayInfo{}).RowsAffected
}
