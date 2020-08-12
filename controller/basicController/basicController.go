package basicController

import (
	"net/http"
	"strconv"

	"github.com/Biubiubiuuuu/yuepuwebsite/service/basicService"

	"github.com/gin-gonic/gin"
)

// @Summary 获取省
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/province [GET]
func QueryProvinces(c *gin.Context) {
	res := basicService.QueryProvinces()
	c.JSON(http.StatusOK, res)
}

// @Summary 根据省代码获取城市
// @tags 基础数据
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param province_code query string true "省代码""
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/city [GET]
func QueryCitysByProvinceCode(c *gin.Context) {
	province_code := c.Query("province_code")
	res := basicService.QueryCitysByProvinceCode(province_code)
	c.JSON(http.StatusOK, res)
}

// @Summary 根据城市代码获取区
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param city_code query string true "城市代码""
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/district [GET]
func QueryDistrictByCityCode(c *gin.Context) {
	city_code := c.Query("city_code")
	res := basicService.QueryDistrictByCityCode(city_code)
	c.JSON(http.StatusOK, res)
}

// @Summary 获取所有区
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/districts [GET]
func QueryDistrict(c *gin.Context) {
	res := basicService.QueryDistrict()
	c.JSON(http.StatusOK, res)
}

// @Summary 根据区代码获取街道
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param district_code query string true "区代码""
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/street [GET]
func QueryStreetByDistrictCode(c *gin.Context) {
	district_code := c.Query("district_code")
	res := basicService.QueryStreetByDistrictCode(district_code)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询面积分类
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/areatype [GET]
func QueryAreaType(c *gin.Context) {
	res := basicService.QueryAreaType()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询租金分类
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/renttype [GET]
func QueryRentType(c *gin.Context) {
	res := basicService.QueryRentType()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询已启用店铺类型
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/enableStoreType [GET]
func QueryEnableStoreType(c *gin.Context) {
	res := basicService.QueryEnableStoreType()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询已启用行业类型
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/enableIndustry [GET]
func QueryEnableIndustry(c *gin.Context) {
	res := basicService.QueryEnableIndustry()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询已启用最上级行业
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/industryParent [GET]
func QueryEnableIndustryByParentID(c *gin.Context) {
	res := basicService.QueryEnableIndustryByParentID()
	c.JSON(http.StatusOK, res)
}

// @Summary 获取所有街道
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/streets [GET]
func QueryStreet(c *gin.Context) {
	res := basicService.QueryStreet()
	c.JSON(http.StatusOK, res)
}

// @Summary 根据上级查询下级行业
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/industrys [GET]
func QueryIndustryByParentID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	res := basicService.QueryIndustryByParentID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 最新动态
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/proInfoDynamic [GET]
func QueryProInfoDynamic(c *gin.Context) {
	res := basicService.QueryProInfoDynamic()
	c.JSON(http.StatusOK, res)
}

// @Summary 获取配置信息
// @tags 基础数据
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/basic/systemConfig [GET]
func QuerySystemConfigByDefault(c *gin.Context) {
	res := basicService.QuerySystemConfigByDefault()
	c.JSON(http.StatusOK, res)
}
