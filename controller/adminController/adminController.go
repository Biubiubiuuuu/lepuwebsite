package adminController

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/adminService"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserLogin true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/login [POST]
func Login(c *gin.Context) {
	req := entity.UserLogin{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = adminService.Login(req, c.ClientIP())
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加面积分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.AreaTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/areaType [POST]
// @Security ApiKeyAuth
func CreateAreaType(c *gin.Context) {
	req := entity.AreaTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = adminService.CreateAreaType(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改面积分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "面积分类ID"
// @Param body body entity.AreaTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/areaType/{id} [PUT]
// @Security ApiKeyAuth
func EditAreaType(c *gin.Context) {
	req := entity.AreaTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditAreaType(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除面积分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "面积分类ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/areaType/{ids} [DELETE]
// @Security ApiKeyAuth
func DelAreaType(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelAreaType(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加租金分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.RentTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rentType [POST]
// @Security ApiKeyAuth
func CreateRentType(c *gin.Context) {
	req := entity.RentTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = adminService.CreateRentType(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改租金分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "租金分类ID"
// @Param body body entity.RentTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rentType/{id} [PUT]
// @Security ApiKeyAuth
func EditRentType(c *gin.Context) {
	req := entity.RentTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditRentType(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除租金分类
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "租金分类ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rentType/{ids} [DELETE]
// @Security ApiKeyAuth
func DelRentType(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelRentType(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加经营业态
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.IndustryRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/industry [POST]
// @Security ApiKeyAuth
func AddIndustry(c *gin.Context) {
	req := entity.IndustryRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = adminService.AddIndustry(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改经营业态
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "经营业态ID"
// @Param body body entity.IndustryRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/industry/{id} [PUT]
// @Security ApiKeyAuth
func EditIndustry(c *gin.Context) {
	req := entity.IndustryRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditIndustry(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除经营业态
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "租金分类ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/industry/{ids} [DELETE]
// @Security ApiKeyAuth
func DelIndustry(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelIndustry(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加店铺类型
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.StoreTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/storeType [POST]
// @Security ApiKeyAuth
func AddStoreType(c *gin.Context) {
	req := entity.StoreTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = adminService.AddStoreType(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改店铺类型
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "经营业态ID"
// @Param body body entity.StoreTypeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/storeType/{id} [PUT]
// @Security ApiKeyAuth
func EditStoreType(c *gin.Context) {
	req := entity.StoreTypeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditStoreType(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除店铺类型
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "租金分类ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/storeType/{ids} [DELETE]
// @Security ApiKeyAuth
func DelStoreType(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelStoreType(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加部门
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.DepartmentRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/department [POST]
// @Security ApiKeyAuth
func AddDepartment(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.DepartmentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddDepartment(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加岗位
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.PostRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/post [POST]
// @Security ApiKeyAuth
func AddPost(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PostRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddPost(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加角色
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.RoleRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/role [POST]
// @Security ApiKeyAuth
func AddRole(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.RoleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddRole(req)
	}
	c.JSON(http.StatusOK, res)
}
