package basicService

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
)

// 查询省
func QueryProvinces() (res entity.ResponseData) {
	data := make(map[string]interface{})
	province := model.Province{}
	provinces := province.QueryProvinces()
	data["provices"] = provinces
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询城市
func QueryCitysByProvinceCode(province_code string) (res entity.ResponseData) {
	data := make(map[string]interface{})
	city := model.City{
		ProvinceCode: province_code,
	}
	citys := city.QueryCitysByProvinceCode()
	data["citys"] = citys
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询区
func QueryDistrictByCityCode(city_code string) (res entity.ResponseData) {
	data := make(map[string]interface{})
	d := model.District{
		CityCode: city_code,
	}
	districts := d.QueryDistrictByCityCode()
	data["districts"] = districts
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询区
func QueryDistrict() (res entity.ResponseData) {
	data := make(map[string]interface{})
	d := model.District{}
	districts := d.QueryDistrict()
	data["districts"] = districts
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询街道
func QueryStreetByDistrictCode(district_code string) (res entity.ResponseData) {
	data := make(map[string]interface{})
	s := model.Street{
		DistrictCode: district_code,
	}
	streets := s.QueryStreetByDistrictCode()
	data["streets"] = streets
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询面积分类
func QueryAreaType() (res entity.ResponseData) {
	data := make(map[string]interface{})
	a := model.AreaType{}
	area_types := a.QueryAreaType()
	data["area_types"] = area_types
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询租金分类
func QueryRentType() (res entity.ResponseData) {
	data := make(map[string]interface{})
	r := model.RentType{}
	rent_types := r.QueryRentType()
	data["rent_types"] = rent_types
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询已启用店铺类型
func QueryEnableStoreType() (res entity.ResponseData) {
	data := make(map[string]interface{})
	s := model.StoreType{}
	store_types := s.QueryEnableStoreType()
	data["store_types"] = store_types
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询已启用行业类型
func QueryEnableIndustry() (res entity.ResponseData) {
	data := make(map[string]interface{})
	i := model.Industry{}
	industrys := i.QueryEnableIndustry()
	data["industrys"] = industrys
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询已启用上级行业
func QueryEnableIndustryByParentID() (res entity.ResponseData) {
	data := make(map[string]interface{})
	industrys := model.QueryEnableIndustryByParentID()
	data["industrys"] = industrys
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}
