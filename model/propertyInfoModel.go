package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
	"strings"
)

// PropertyInfo 物业信息
type PropertyInfo struct {
	Model
	IndustryID     int64           `json:"industry_id"`                                                                // 经营业态ID
	Title          string          `gorm:"not null;size:60;" json:"title"`                                             // 标题
	Nickname       string          `gorm:"size:50;" json:"nickname"`                                                   // 联系人
	Telephone      string          `gorm:"not null;unique;size:30;" json:"telephone"`                                  // 联系手机
	ShopName       string          `gorm:"size:100;" json:"shop_name"`                                                 // 店名（后台录入）
	Image          string          `gorm:"size:200;" json:"image"`                                                     // 图片
	Video          string          `json:"video"`                                                                      // 视频（后台录入）
	BusType        string          `gorm:"size:1;" json:"bus_type"`                                                    // 业务类型（后台录入）0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库
	ModelType      string          `gorm:"size:1;" json:"model_type"`                                                  // 模型类型（后台录入）0-转让 ｜ 1-出售 ｜ 3-出租 | 4-求租 ｜ 5-求购
	ProvinceCode   string          `gorm:"not null;size:10;" json:"province_code"`                                     // 省代码
	CityCode       string          `gorm:"not null;size:10;" json:"city_code"`                                         // 城市代码
	DistrictCode   string          `gorm:"not null;size:10;" json:"district_code"`                                     // 区代码
	StreetCode     string          `gorm:"not null;size:10;" json:"street_code"`                                       // 街道代码
	Address        string          `json:"address"`                                                                    // 详细地址
	StoreTypeID    int64           `json:"store_type_id"`                                                              // 店铺类型ID
	AreaTypeID     int64           `json:"area_type_id"`                                                               // 面积分类ID （后台录入或者自动判断）
	RentTypeID     int64           `json:"rent_type_id"`                                                               // 租金分类ID （后台录入或者自动判断）
	MinArea        float64         `json:"min_area"`                                                                   // 最小面积（单位：平方米）
	MaxArea        float64         `json:"max_area"`                                                                   // 最大面积（单位：平方米）
	MinRent        float64         `json:"min_rent"`                                                                   // 最低租金（单位：元/月）
	MaxRent        float64         `json:"max_rent"`                                                                   // 最高租金（单位：元/月）
	Lots           []Lot           `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"lots"`            // 考虑地段
	Idling         bool            `json:"idling"`                                                                     // 可否空转
	InOperation    string          `gorm:"size:2;" json:"in_operation"`                                                // 是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中
	Area           float64         `json:"area"`                                                                       // 面积（单位：平方米）
	Rent           float64         `json:"rent"`                                                                       // 租金（单位：元/月）
	TransferFee    float64         `json:"transfer_fee"`                                                               // 转让费用（单位：万元 不录入则前台显示为面议）
	IndustryRanges []IndustryRange `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"industry_ranges"` // 适合经营范围
	Description    string          `json:"description"`                                                                // 详细描述
	ExplicitTel    bool            `json:"explicit_tel"`                                                               // 是否外显号码 true：客户号码 ｜ false：发布者号码
	Tel1           string          `gorm:"size:30;" json:"tel1"`                                                       // 外显号码1
	Tel2           string          `gorm:"size:30;" json:"tel2"`                                                       // 外显号码2
	Audit          bool            `json:"audit"`                                                                      // 是否审核 true：已审核 ｜ false：待审核 （后台录入）
	AuditID        int64           `json:"audit_id"`                                                                   // 审核人ID （后台录入）
	Protect        bool            `json:"protect"`                                                                    // 是否保护 true：已保护 ｜ false：未保护 （后台录入）
	QuotedPrice    float64         `json:"quoted_price"`                                                               // 报价（后台录入，保护时显示）
	Pictures       []Picture       `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"pictures"`        // 店图集（后台录入）
	Status         bool            `json:"status"`                                                                     // 是否成功 true：已成功 ｜ false：未成功 （后台录入）
	SourceID       int64           `json:"source_id"`                                                                  // 来源ID
	SourceInfo     string          `gorm:"size:200" json:"source_info"`                                                // 来源描述
	Remake         string          `gorm:"size:200" json:"remake"`                                                     // 跟进备注
}

// PropertyInfoScan 物业信息详细
type PropertyInfoScan struct {
	PropertyInfo
	IndustryName  string `json:"industry_name"`   // 行业名称
	ProvinceName  string `json:"province_name"`   // 省名称
	CityName      string `json:"city_name"`       // 城市名称
	DistrictName  string `json:"district_name"`   // 区名称
	StreetName    string `json:"street_name"`     // 街道名称
	StoreTypeName string `json:"store_type_name"` // 店铺类型名称
	AreaType_name string `json:"area_type_name"`  // 面积分类名称
	RentTypeName  string `json:"rent_type_name"`  // 租金分类名称
	AuditName     string `json:"audit_name"`      // 审核人
	SourceName    string `json:"source_name"`     // 来源人
}

// IndustryRange 适合经营范围
type IndustryRange struct {
	ID             int64
	IndustryID     int64  `json:"industry_id"`                   // 行业ID
	IndustryName   string `json:"industry_name"`                 // 行业名称
	PropertyInfoID int64  `gorm:"INDEX" json:"property_info_id"` // 物业信息ID
}

// 图片
type Picture struct {
	ID             int64
	Url            string `json:"url"`                           // 店铺图
	PropertyInfoID int64  `gorm:"INDEX" json:"property_info_id"` // 物业信息ID
}

type Lot struct {
	ID             int64
	DistrictCode   string `json:"district_code"`                 // 区代码
	DistrictName   string `json:"district_name"`                 // 区名
	PropertyInfoID int64  `gorm:"INDEX" json:"property_info_id"` // 物业信息ID
}

// 添加物业信息
func (p *PropertyInfo) CreatePropertyInfo() error {
	db := mysql.GetMysqlDB()
	return db.Create(p).Error
}

// 查询物业信息 by id
func (p *PropertyInfoScan) QueryPropertyInfoByID() error {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges").Preload("Lots").Preload("Pictures")
	query = query.Select("property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id AND property_info.source_id = user.id")
	return query.First(&p, p.ID).Error
}

// 查询物业信息 by source_id
func (p *PropertyInfo) QueryPropertyInfoByUserID() (propertyInfoScans []PropertyInfoScan) {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges").Preload("Lots").Preload("Pictures")
	query = query.Select("property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id AND property_info.source_id = user.id")
	query = query.Where("property_info.source_id = ?", p.SourceID)
	query.Find(&propertyInfoScans)
	return
}

// 修改物业信息 by id
func (p *PropertyInfo) EditPropertyInfoByID(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	db.Model(&p).Association("IndustryRanges").Replace(p.IndustryRanges)
	db.Association("Pictures").Replace(p.Pictures)
	db.Association("Lots").Replace(p.Lots)
	return db.Model(&p).Update(args).Error
}

// 查看物业信息
func (p *PropertyInfo) QueryPropertyInfo(pageSize int, page int, args map[string]interface{}) (propertyInfoScans []PropertyInfoScan, count int) {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges")
	query = query.Select("property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id AND property_info.source_id = user.id")
	if v, ok := args["telephone"]; ok {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("property_info.telephone like ?", buf.String())
	}
	if v, ok := args["title"]; ok {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("property_info.title like ?", buf.String())
	}
	if v, ok := args["province_code"]; ok {
		query = query.Where("province.code = ?", v.(string))
	}
	if v, ok := args["city_code"]; ok {
		query = query.Where("city.code = ?", v.(string))
	}
	if v, ok := args["district_code"]; ok {
		query = query.Where("district.code = ?", v.(string))
	}
	if v, ok := args["street_code"]; ok {
		query = query.Where("street.code = ?", v.(string))
	}
	if v, ok := args["audit"]; ok {
		query = query.Where("property_info.audit = ?", v.(string))
	}
	if v, ok := args["industry_id"]; ok {
		query = query.Where("property_info.industry_id = ?", v.(string))
	}
	if v, ok := args["area_type_id"]; ok {
		query = query.Where("property_info.area_type_id = ?", v.(string))
	}
	if v, ok := args["rent_type_id"]; ok {
		query = query.Where("property_info.rent_type_id = ?", v.(string))
	}
	if v, ok := args["min_area"]; ok {
		query = query.Where("property_info.area BETWEEN ? AND ?", v.(string), v.(string))
	}
	v1, ok1 := args["min_area"]
	v2, ok2 := args["max_area"]
	if ok1 && ok2 {
		query = query.Where("property_info.area BETWEEN ? AND ?", v1.(string), v2.(string))
	} else if ok1 {
		query = query.Where("property_info.area >= ?", v1.(string))
	} else if ok2 {
		query = query.Where("property_info.area <= ?", v2.(string))
	}
	v3, ok3 := args["min_rent"]
	v4, ok4 := args["max_rent"]
	if ok3 && ok4 {
		query = query.Where("property_info.rent BETWEEN ? AND ?", v3.(string), v4.(string))
	} else if ok3 {
		query = query.Where("property_info.rent >= ?", v3.(string))
	} else if ok4 {
		query = query.Where("property_info.rent <= ?", v4.(string))
	}
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&propertyInfoScans)
	return
}
