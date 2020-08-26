package adminController

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/controller/commonController"
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/configHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/fileHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/adminService"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary 查询面积分类详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "面积分类ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/areaType/{id} [GET]
// @Security ApiKeyAuth
func QueryAreaTypeInfoById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryAreaTypeInfoById(id)
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

// @Summary 查询租金分类详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "租金分类ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rentType/{id} [GET]
// @Security ApiKeyAuth
func QueryRentTypeInfoById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryRentTypeInfoById(id)
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

// @Summary 查询部门
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param name query string false "部门名称"
// @Param enable query string false "是否启用"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/department [GET]
// @Security ApiKeyAuth
func QueryDepartment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryDepartment(pageSize, page, c.Query("name"), c.Query("enable"))
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

// @Summary 修改部门
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "部门ID"
// @Param body body entity.DepartmentRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/department/{id} [PUT]
// @Security ApiKeyAuth
func EditDepartment(c *gin.Context) {
	req := entity.DepartmentRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditDepartment(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除部门
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "部门ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/department/{ids} [DELETE]
// @Security ApiKeyAuth
func DelDepartment(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelDepartment(ids)
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

// @Summary 修改岗位
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "岗位ID"
// @Param body body entity.PostRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/post/{id} [PUT]
// @Security ApiKeyAuth
func EditPost(c *gin.Context) {
	req := entity.PostRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditPost(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除岗位
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "岗位ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/post/{ids} [DELETE]
// @Security ApiKeyAuth
func DelPost(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelPost(ids)
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

// @Summary 修改角色
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "角色ID"
// @Param body body entity.RoleRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/role/{id} [PUT]
// @Security ApiKeyAuth
func EditRole(c *gin.Context) {
	req := entity.RoleRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditRole(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除角色
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "角色ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/role/{ids} [DELETE]
// @Security ApiKeyAuth
func DelRole(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelRole(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询店铺类型详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "店铺类型ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/storeType/{id} [GET]
// @Security ApiKeyAuth
func QueryStoreTypeByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryStoreTypeByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询岗位详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "岗位ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/post/{id} [GET]
// @Security ApiKeyAuth
func QueryPostByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryPostByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询部门详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "部门ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/department/{id} [GET]
// @Security ApiKeyAuth
func QueryDepartmentByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryDepartmentByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询所有岗位
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param name query string false "岗位名称"
// @Param code query string false "岗位编码"
// @Param enable query string false "是否启用"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/post [GET]
// @Security ApiKeyAuth
func QueryPost(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryPost(pageSize, page, c.Query("name"), c.Query("code"), c.Query("enable"))
	c.JSON(http.StatusOK, res)
}

// @Summary 查询店铺类型
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param name query string false "店铺类型名称"
// @Param enable query string false "是否启用"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/storeType [GET]
// @Security ApiKeyAuth
func QueryStoreType(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryStoreType(pageSize, page, c.Query("name"), c.Query("enable"))
	c.JSON(http.StatusOK, res)
}

// @Summary 查询行业分类详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "行业分类ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/industry/{id} [GET]
// @Security ApiKeyAuth
func QueryIndustryByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryIndustryByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询行业类型
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param name query string false "行业名称"
// @Param enable query string false "是否启用"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/industry [GET]
// @Security ApiKeyAuth
func QueryIndustry(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryIndustry(pageSize, page, c.Query("name"), c.Query("enable"))
	c.JSON(http.StatusOK, res)
}

// @Summary 查询面积分类
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/areaType [GET]
// @Security ApiKeyAuth
func QueryAreaType(c *gin.Context) {
	res := adminService.QueryAreaType()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询租金分类
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rentType [GET]
// @Security ApiKeyAuth
func QueryRentType(c *gin.Context) {
	res := adminService.QueryRentType()
	c.JSON(http.StatusOK, res)
}

// @Summary 查询角色
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param name query string false "角色名称"
// @Param enable query string false "是否启用"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/role [GET]
// @Security ApiKeyAuth
func QueryRole(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryRole(pageSize, page, c.Query("name"), c.Query("enable"))
	c.JSON(http.StatusOK, res)
}

// @Summary 查询角色详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "角色ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/role/{id} [GET]
// @Security ApiKeyAuth
func QueryRoleByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryRoleByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加菜单
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.MenuRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/menu [POST]
// @Security ApiKeyAuth
func AddMenu(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.MenuRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddMenu(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改菜单
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "菜单ID"
// @Param body body entity.MenuRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/menu/{id} [PUT]
// @Security ApiKeyAuth
func Editmenu(c *gin.Context) {
	req := entity.MenuRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditMenu(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除菜单
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "菜单ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/menu/{ids} [DELETE]
// @Security ApiKeyAuth
func DelMenu(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelMenus(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 查看菜单详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "菜单ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/menu/{id} [GET]
// @Security ApiKeyAuth
func QueryMenuByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryMenuByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 查看菜单
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/menu [GET]
// @Security ApiKeyAuth
func QueryMenu(c *gin.Context) {
	res := adminService.QueryMenu()
	c.JSON(http.StatusOK, res)
}

// @Summary 获取员工详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "员工ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/employee/{id} [GET]
// @Security ApiKeyAuth
func QueryEmployeeById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryEmployeeById(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改员工信息
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "员工ID"
// @Param body body entity.EditEmployeeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/employee/{id} [PUT]
// @Security ApiKeyAuth
func EditUser(c *gin.Context) {
	req := entity.EditEmployeeRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditAdminToUser(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改用户密码
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.EditUserPass true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/edituserpass [POST]
// @Security ApiKeyAuth
func EditUserPass(c *gin.Context) {
	req := entity.EditUserPass{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		var token string
		if token, res = commonController.GetToken(c); res.Status {
			res = userService.EditUserPass(token, req)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 查看用户信息
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin [GET]
// @Security ApiKeyAuth
func QueryUser(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
		res = userService.QueryUserByToken(token)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加员工
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param body body entity.AddEmployeeRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/employee [POST]
// @Security ApiKeyAuth
func AddEmployee(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.AddEmployeeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddEmployee(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除员工
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "用户ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/employee/{ids} [DELETE]
// @Security ApiKeyAuth
func DelEmployee(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelEmployee(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 获取员工列表
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param nickname query string false "用户昵称"
// @Param username query string false "用户名"
// @Param enable query string false "是否启用"
// @Param telephone query string false "电话号码"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/employee [GET]
// @Security ApiKeyAuth
func GetEmployee(c *gin.Context) {
	args := map[string]interface{}{
		"nickname":  c.Query("nickname"),
		"username":  c.Query("username"),
		"enable":    c.Query("enable"),
		"telephone": c.Query("telephone"),
	}
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryAllUser(pageSize, page, args)
	c.JSON(http.StatusOK, res)
}

// @Summary 查询我的已发布物业信息
// @tags 后台
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param industry_id query string false "行业ID"
// @Param store_type_id query string false "店铺类型ID"
// @Param province_code query string false "省代码"
// @Param city_code query string false "城市代码"
// @Param district_code query string false "区代码"
// @Param street_code query string false "街道代码"
// @Param area_type_id query string false "面积范围ID"
// @Param rent_type_id query string false "租金范围ID"
// @Param min_area query string false "最小面积"
// @Param max_area query string false "最大面积"
// @Param min_rent query string false "最小租金"
// @Param max_rent query string false "最大租金"
// @Param model_type query string true "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租 | 4-求租 ｜ 5-求购"
// @Param bus_type query string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param sort_condition query string false "排序 area-面积 ｜ rent-租金 ｜ created_at-发布时间（默认）"
// @Param status query string false "是否成功 查询我的已成功信息传true；查询我的历史重点信息true"
// @Param protect query string false "是否保护 查询我的重点信息传true；查询我的历史重点信息true"
// @Param created_at query string false "创建时间"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfo [GET]
// @Security ApiKeyAuth
func SearchPropertyInfo(c *gin.Context) {
	args := map[string]interface{}{
		"store_type_id":  c.Query("store_type_id"),
		"industry_id":    c.Query("industry_id"),
		"province_code":  c.Query("province_code"),
		"city_code":      c.Query("city_code"),
		"district_code":  c.Query("district_code"),
		"street_code":    c.Query("street_code"),
		"area_type_id":   c.Query("area_type_id"),
		"min_area":       c.Query("min_area"),
		"max_area":       c.Query("max_area"),
		"min_rent":       c.Query("min_rent"),
		"max_rent":       c.Query("max_rent"),
		"rent_type_id":   c.Query("rent_type_id"),
		"model_type":     c.Query("model_type"),
		"bus_type":       c.Query("bus_type"),
		"sort_condition": c.DefaultQuery("sort_condition", "created_at"),
	}
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	var token string
	res := entity.ResponseData{}
	if token, res = commonController.GetToken(c); res.Status {
		res = adminService.SearchPropertyInfo(token, pageSize, page, args)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 查看物业详情
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfo/{id} [GET]
// @Security ApiKeyAuth
func QueryPropertyInfoByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := userService.QueryPropertyInfoByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改物业信息
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "物业信息ID"
// @Param industry_id formData string true "经营业态ID"
// @Param title formData string true "标题"
// @Param nickname formData string true "联系人"
// @Param telephone formData string true "联系手机"
// @Param image formData file false "图片"
// @Param province_code formData string true "省代码""
// @Param city_code formData string true "城市代码"
// @Param district_code formData string true "区代码"
// @Param street_code formData string false "街道代码"
// @Param address formData string true "详细地址"
// @Param store_type_id formData string true "店铺类型ID"
// @Param idling formData bool true "可否空转"
// @Param area formData string true "面积（单位：平方米）"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param rent formData string true "租金（单位：元/月）"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param description formData string false "详细描述""
// @Param explicit_tel formData bool false "是否外显号码"
// @Param tel1 formData string false "外显号码1"
// @Param tel2 formData string false "外显号码2"
// @Param quoted_price formData string false "报价"
// @Param remake formData string false "跟进备注"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfo/{id} [PUT]
// @Security ApiKeyAuth
func EditUserStoretransfer(c *gin.Context) {
	res := entity.ResponseData{}
	industryID, _ := strconv.ParseInt(c.PostForm("industry_id"), 10, 64)
	storeTypeID, _ := strconv.ParseInt(c.PostForm("store_type_id"), 10, 64)
	idling, _ := strconv.ParseBool(c.PostForm("idling"))
	area, _ := strconv.ParseFloat(c.PostForm("area"), 64)
	rent, _ := strconv.ParseFloat(c.PostForm("rent"), 64)
	transferFee, _ := strconv.ParseFloat(c.PostForm("transfer_fee"), 64)
	industryRangesArr := strings.Split(c.PostForm("industry_ranges"), ",")
	var industryRanges []int64
	for _, item := range industryRangesArr {
		id, _ := strconv.ParseInt(item, 10, 64)
		industryRanges = append(industryRanges, id)
	}
	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	quoted_price, _ := strconv.ParseFloat(c.PostForm("quoted_price"), 64)

	explicit_tel, _ := strconv.ParseBool(c.PostForm("explicit_tel"))
	req := entity.AdminStoretransferRequest{
		IndustryID:     industryID,
		Title:          c.PostForm("title"),
		Nickname:       c.PostForm("nickname"),
		Telephone:      c.PostForm("telephone"),
		Image:          image,
		ProvinceCode:   c.PostForm("province_code"),
		CityCode:       c.PostForm("city_code"),
		DistrictCode:   c.PostForm("district_code"),
		StreetCode:     c.PostForm("stree_code"),
		Address:        c.PostForm("address"),
		StoreTypeID:    storeTypeID,
		Idling:         idling,
		Area:           area,
		Rent:           rent,
		TransferFee:    transferFee,
		IndustryRanges: industryRanges,
		Description:    c.PostForm("description"),
		InOperation:    c.PostForm("in_operation"),
		ExplicitTel:    explicit_tel,
		Tel1:           c.PostForm("tel1"),
		Tel2:           c.PostForm("tel2"),
		QuotedPrice:    quoted_price,
		Remake:         c.PostForm("remake"),
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res = adminService.EditUserStoretransfer(id, req)
	c.JSON(http.StatusOK, res)
}

// @Summary 保护
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfo/protect/{id} [POST]
// @Security ApiKeyAuth
func EditProtectionProInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.EditProtectionProInfo(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 取消保护
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfo/notprotect/{id} [POST]
// @Security ApiKeyAuth
func EditNotProtectionProInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.EditNotProtectionProInfo(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 上传图集图片（单张）
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "物业信息ID"
// @Param image formData file false "图片"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfos/{id}/picture [POST]
// @Security ApiKeyAuth
func AddPictures(c *gin.Context) {
	res := entity.ResponseData{}
	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}
	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		res.Message = "图片下载失败"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.AddPictures(id, image)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除图片（单张）
// @tags 后台
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pro_id path string true "物业信息ID"
// @Param pri_id path string true "图片ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/propertyInfos/{pro_id}/picture/{pri_id} [DELETE]
// @Security ApiKeyAuth
func DelPrictures(c *gin.Context) {
	pro_id, _ := strconv.ParseInt(c.Param("pro_id"), 10, 64)
	pri_id, _ := strconv.ParseInt(c.Param("pri_id"), 10, 64)
	res := adminService.DelPrictures(pro_id, pri_id)
	c.JSON(http.StatusOK, res)
}

// @Summary 物业添加跟单记录
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Param body body entity.AddProLog true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfo/log/{id} [POST]
// @Security ApiKeyAuth
func AddProInfoLog(c *gin.Context) {
	var token string
	res := entity.ResponseData{}
	req := entity.AddProLog{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		if token, res = commonController.GetToken(c); res.Status {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			res = adminService.AddProInfoLog(token, id, req)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 跟单记录详情
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Param body body entity.AddProLog true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfo/log/{id} [GET]
// @Security ApiKeyAuth
func QueryProInfoLog(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryByProInfoID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 发布物业信息
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param title formData string true "标题"
// @Param nickname formData string true "客户姓名"
// @Param telephone formData string true "联系手机"
// @Param store_name formData string true "店名"
// @Param image formData file false "图片"
// @Param bus_type formData string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param model_type formData string false "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租"
// @Param province_code formData string true "省代码""
// @Param city_code formData string true "城市代码"
// @Param district_code formData string true "区代码"
// @Param street_code formData string false "街道代码"
// @Param address formData string true "详细地址"
// @Param industry_id formData int true "经营业态ID"
// @Param store_type_id formData int true "店铺类型ID"
// @Param idling formData bool true "可否空转"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param area formData string true "面积（单位：平方米）"
// @Param rent formData string true "租金（单位：元/月）"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param description formData string false "详细描述"
// @Param explicit_tel formData bool false "是否外显号码"
// @Param tel1 formData string false "外显号码1"
// @Param tel2 formData string false "外显号码2"
// @Param quoted_price formData string false "报价"
// @Param remake formData string false "跟进备注"
// @Param protect formData bool false "是否保护"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/new/propertyInfo [POST]
// @Security ApiKeyAuth
func AddProInfo(c *gin.Context) {
	var token string
	res := entity.ResponseData{}
	industryID, _ := strconv.ParseInt(c.PostForm("industry_id"), 10, 64)
	storeTypeID, _ := strconv.ParseInt(c.PostForm("store_type_id"), 10, 64)
	idling, _ := strconv.ParseBool(c.PostForm("idling"))

	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	explicit_tel, _ := strconv.ParseBool(c.PostForm("explicit_tel"))
	protect, _ := strconv.ParseBool(c.PostForm("protect"))
	req := entity.AddPropertyInfoRequest{
		Title:          c.PostForm("title"),
		Nickname:       c.PostForm("nickname"),
		Telephone:      c.PostForm("telephone"),
		Image:          image,
		ProvinceCode:   c.PostForm("province_code"),
		CityCode:       c.PostForm("city_code"),
		DistrictCode:   c.PostForm("district_code"),
		StreetCode:     c.PostForm("street_code"),
		Address:        c.PostForm("address"),
		StoreTypeID:    storeTypeID,
		Idling:         idling,
		InOperation:    c.PostForm("in_operation"),
		Area:           c.PostForm("area"),
		Rent:           c.PostForm("rent"),
		TransferFee:    c.PostForm("transfer_fee"),
		IndustryRanges: c.PostForm("industry_ranges"),
		Description:    c.PostForm("description"),
		IndustryID:     industryID,
		QuotedPrice:    c.PostForm("quoted_price"),
		ExplicitTel:    explicit_tel,
		Remake:         c.PostForm("remake"),
		ShopName:       c.PostForm("shop_name"),
		Tel1:           c.PostForm("tel1"),
		Tel2:           c.PostForm("tel2"),
		Protect:        protect,
		BusType:        c.PostForm("bus_type"),
		ModelType:      c.PostForm("model_type"),
	}
	if token, res = commonController.GetToken(c); res.Status {
		res = adminService.AddProInfo(token, req)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改物业信息
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string false "ID"
// @Param title formData string true "标题"
// @Param nickname formData string true "客户姓名"
// @Param telephone formData string true "联系手机"
// @Param store_name formData string true "店名"
// @Param image formData file false "图片"
// @Param bus_type formData string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param model_type formData string false "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租"
// @Param province_code formData string true "省代码""
// @Param city_code formData string true "城市代码"
// @Param district_code formData string true "区代码"
// @Param street_code formData string false "街道代码"
// @Param address formData string true "详细地址"
// @Param industry_id formData int true "经营业态ID"
// @Param store_type_id formData int true "店铺类型ID"
// @Param idling formData bool true "可否空转"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param area formData string true "面积（单位：平方米）"
// @Param rent formData string true "租金（单位：元/月）"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param description formData string false "详细描述"
// @Param explicit_tel formData bool false "是否外显号码"
// @Param tel1 formData string false "外显号码1"
// @Param tel2 formData string false "外显号码2"
// @Param quoted_price formData string false "报价"
// @Param remake formData string false "跟进备注"
// @Param protect formData bool false "是否保护"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/new/propertyInfo/{id} [PUT]
// @Security ApiKeyAuth
func EditProInfo(c *gin.Context) {
	res := entity.ResponseData{}
	industryID, _ := strconv.ParseInt(c.PostForm("industry_id"), 10, 64)
	storeTypeID, _ := strconv.ParseInt(c.PostForm("store_type_id"), 10, 64)
	idling, _ := strconv.ParseBool(c.PostForm("idling"))

	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	explicit_tel, _ := strconv.ParseBool(c.PostForm("explicit_tel"))
	protect, _ := strconv.ParseBool(c.PostForm("protect"))
	req := entity.AddPropertyInfoRequest{
		Title:          c.PostForm("title"),
		Nickname:       c.PostForm("nickname"),
		Telephone:      c.PostForm("telephone"),
		Image:          image,
		ProvinceCode:   c.PostForm("province_code"),
		CityCode:       c.PostForm("city_code"),
		DistrictCode:   c.PostForm("district_code"),
		StreetCode:     c.PostForm("street_code"),
		Address:        c.PostForm("address"),
		StoreTypeID:    storeTypeID,
		Idling:         idling,
		InOperation:    c.PostForm("in_operation"),
		Area:           c.PostForm("area"),
		Rent:           c.PostForm("rent"),
		TransferFee:    c.PostForm("transfer_fee"),
		IndustryRanges: c.PostForm("industry_ranges"),
		Description:    c.PostForm("description"),
		IndustryID:     industryID,
		QuotedPrice:    c.PostForm("quoted_price"),
		ExplicitTel:    explicit_tel,
		Remake:         c.PostForm("remake"),
		ShopName:       c.PostForm("shop_name"),
		Tel1:           c.PostForm("tel1"),
		Tel2:           c.PostForm("tel2"),
		Protect:        protect,
		BusType:        c.PostForm("bus_type"),
		ModelType:      c.PostForm("model_type"),
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res = adminService.EditProInfo(id, req)
	c.JSON(http.StatusOK, res)
}

// @Summary 留言列表
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/leaveMessage [GET]
// @Security ApiKeyAuth
func QueryLeaveMessage(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryLeaveMessage(pageSize, page)
	c.JSON(http.StatusOK, res)
}

// @Summary 删除物业信息
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id path string false "物业ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/new/propertyInfo/{id} [DELETE]
// @Security ApiKeyAuth
func DelProInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.DelProInfo(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 留言详情
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id path string false "留言ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/leaveMessage/{id} [GET]
// @Security ApiKeyAuth
func QueryLeaveMessageByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryLeaveMessageByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 举报列表
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/report [GET]
// @Security ApiKeyAuth
func QueryReport(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryReport(pageSize, page)
	c.JSON(http.StatusOK, res)
}

// @Summary 举报详情
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id path string false "举报ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/report/{id} [GET]
// @Security ApiKeyAuth
func QueryReportByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryReportByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加轮播
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param url formData file false "图片"
// @Param link formData string false "跳转链接"
// @Param sort formData int false "排序"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/carousel [POST]
// @Security ApiKeyAuth
func AddCarousel(c *gin.Context) {
	res := entity.ResponseData{}
	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	req := entity.CarouselRequest{
		Url:  image,
		Link: c.PostForm("link"),
		Sort: sort,
	}
	res = adminService.AddCarousel(req)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改轮播
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string false "轮播ID"
// @Param body body entity.CarouselRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/carousel/{id} [PUT]
// @Security ApiKeyAuth
func EditCarousel(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.CarouselRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditCarousel(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除轮播
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/carousel/{ids} [DELETE]
// @Security ApiKeyAuth
func DelCarousel(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelCarousel(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 轮播详情
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id path string false "轮播ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/carousel/{id} [GET]
// @Security ApiKeyAuth
func QueryCarouselByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryCarouselByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 轮播列表
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/carousel [GET]
// @Security ApiKeyAuth
func QueryCarousel(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryCarousel(pageSize, page)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加广告
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param body body entity.AdvertRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/advert [POST]
// @Security ApiKeyAuth
func AddAdvert(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.AdvertRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddAdvert(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改广告
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "广告ID"
// @Param body body entity.AdvertRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/advert/{id} [PUT]
// @Security ApiKeyAuth
func EditAdvert(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.AdvertRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditAdvert(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除广告
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/advert/{ids} [DELETE]
// @Security ApiKeyAuth
func DelAdvert(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelAdvert(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 广告详情
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param id path string false "广告ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/advert/{id} [GET]
// @Security ApiKeyAuth
func QueryAdvertByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryAdvertByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 广告列表
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param hot query bool false "首页最热推广"
// @Param floor query bool false "F楼"
// @Param type query string false "信息列表推广 1-一栏四分之一图片广告 | 2-二栏四分之一图片广告 | 3-三栏重点推荐 | 4-五栏框架广告"
// @Param enable query bool false "是否已审核"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/advert [GET]
// @Security ApiKeyAuth
func QueryAdvert(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := adminService.QueryAdvert(pageSize, page, map[string]interface{}{
		"hot":    c.Query("hot"),
		"floor":  c.Query("floor"),
		"type":   c.Query("type"),
		"enable": c.Query("enable"),
	})
	c.JSON(http.StatusOK, res)
}

// @Summary 发布求租求购
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param title formData string true "标题"
// @Param nickname formData string true "客户姓名"
// @Param telephone formData string true "联系手机"
// @Param store_name formData string true "店名"
// @Param image formData file false "图片"
// @Param bus_type formData string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param model_type formData string false "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租"
// @Param city_code formData string true "城市代码"
// @Param industry_id formData int true "经营业态ID"
// @Param min_area query string false "最小面积"
// @Param max_area query string false "最大面积"
// @Param min_rent query string false "最小租金"
// @Param max_rent query string false "最大租金"
// @Param idling formData bool true "可否空转"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param description formData string false "详细描述"
// @Param explicit_tel formData bool false "是否外显号码"
// @Param tel1 formData string false "外显号码1"
// @Param tel2 formData string false "外显号码2"
// @Param quoted_price formData string false "报价"
// @Param remake formData string false "跟进备注"
// @Param protect formData bool false "是否保护"
// @Param source_info formData string false "来源描述"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/new/qzqgPropertyInfo [POST]
// @Security ApiKeyAuth
func AddQZQGProInfo(c *gin.Context) {
	var token string
	res := entity.ResponseData{}
	industryID, _ := strconv.ParseInt(c.PostForm("industry_id"), 10, 64)
	idling, _ := strconv.ParseBool(c.PostForm("idling"))

	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	explicit_tel, _ := strconv.ParseBool(c.PostForm("explicit_tel"))
	protect, _ := strconv.ParseBool(c.PostForm("protect"))
	min_area, _ := strconv.ParseFloat(c.PostForm("min_area"), 64)
	max_area, _ := strconv.ParseFloat(c.PostForm("max_area"), 64)
	min_rent, _ := strconv.ParseFloat(c.PostForm("min_rent"), 64)
	max_rent, _ := strconv.ParseFloat(c.PostForm("max_rent"), 64)
	req := entity.AddQZQGPropertyInfoRequest{
		Title:       c.PostForm("title"),
		Nickname:    c.PostForm("nickname"),
		Telephone:   c.PostForm("telephone"),
		Image:       image,
		CityCode:    c.PostForm("city_code"),
		Idling:      idling,
		InOperation: c.PostForm("in_operation"),
		TransferFee: c.PostForm("transfer_fee"),
		Description: c.PostForm("description"),
		IndustryID:  industryID,
		QuotedPrice: c.PostForm("quoted_price"),
		ExplicitTel: explicit_tel,
		Remake:      c.PostForm("remake"),
		ShopName:    c.PostForm("shop_name"),
		Tel1:        c.PostForm("tel1"),
		Tel2:        c.PostForm("tel2"),
		Protect:     protect,
		BusType:     c.PostForm("bus_type"),
		ModelType:   c.PostForm("model_type"),
		MinArea:     min_area,
		MaxArea:     max_area,
		MinRent:     min_rent,
		MaxRent:     max_rent,
		Lots:        c.PostForm("lots"),
		SourceInfo:  c.PostForm("source_info"),
	}
	if token, res = commonController.GetToken(c); res.Status {
		res = adminService.AddQZQGProInfo(token, req)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改求租求购
// @tags 后台
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string false "ID"
// @Param title formData string true "标题"
// @Param nickname formData string true "客户姓名"
// @Param telephone formData string true "联系手机"
// @Param store_name formData string true "店名"
// @Param image formData file false "图片"
// @Param bus_type formData string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param model_type formData string false "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租"
// @Param city_code formData string true "城市代码"
// @Param industry_id formData int true "经营业态ID"
// @Param min_area query string false "最小面积"
// @Param max_area query string false "最大面积"
// @Param min_rent query string false "最小租金"
// @Param max_rent query string false "最大租金"
// @Param idling formData bool true "可否空转"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param description formData string false "详细描述"
// @Param explicit_tel formData bool false "是否外显号码"
// @Param tel1 formData string false "外显号码1"
// @Param tel2 formData string false "外显号码2"
// @Param quoted_price formData string false "报价"
// @Param remake formData string false "跟进备注"
// @Param protect formData bool false "是否保护"
// @Param source_info formData string false "来源描述"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/new/qzqgPropertyInfo/{id} [PUT]
// @Security ApiKeyAuth
func EditQZQGProInfo(c *gin.Context) {
	res := entity.ResponseData{}
	industryID, _ := strconv.ParseInt(c.PostForm("industry_id"), 10, 64)
	idling, _ := strconv.ParseBool(c.PostForm("idling"))
	// 获取主机头
	r := c.Request
	host := r.Host
	if strings.HasPrefix(host, "http://") == false {
		host = "http://" + host
	}

	var image string
	if file, err := c.FormFile("image"); err == nil {
		// 文件名 避免重复取uuid
		var filename string
		uuid, _ := uuid.NewUUID()
		arr := strings.Split(file.Filename, ".")
		if strings.EqualFold(arr[len(arr)-1], "png") {
			filename = uuid.String() + ".png"
		} else if strings.EqualFold(arr[len(arr)-1], "jpg") {
			filename = uuid.String() + ".jpg"
		} else if strings.EqualFold(arr[len(arr)-1], "jpeg") {
			filename = uuid.String() + ".jpeg"
		} else if strings.EqualFold(arr[len(arr)-1], "gif") {
			filename = uuid.String() + ".gif"
		} else {
			res.Message = "图片格式只支持png、jpg、jpeg、gif"
			c.JSON(http.StatusOK, res)
			return
		}
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + filename
		if err := c.SaveUploadedFile(file, pathFile); err == nil {
			image = host + "/" + pathFile
		}
	}
	if image == "" {
		pathFile := configHelper.ImageDir
		if !fileHelper.IsExist(pathFile) {
			fileHelper.CreateDir(pathFile)
		}
		pathFile = pathFile + "default-store.7e8792da.jpg"
		image = host + "/" + pathFile
	}
	explicit_tel, _ := strconv.ParseBool(c.PostForm("explicit_tel"))
	protect, _ := strconv.ParseBool(c.PostForm("protect"))
	min_area, _ := strconv.ParseFloat(c.PostForm("min_area"), 64)
	max_area, _ := strconv.ParseFloat(c.PostForm("max_area"), 64)
	min_rent, _ := strconv.ParseFloat(c.PostForm("min_rent"), 64)
	max_rent, _ := strconv.ParseFloat(c.PostForm("max_rent"), 64)
	req := entity.AddQZQGPropertyInfoRequest{
		Title:       c.PostForm("title"),
		Nickname:    c.PostForm("nickname"),
		Telephone:   c.PostForm("telephone"),
		Image:       image,
		CityCode:    c.PostForm("city_code"),
		Idling:      idling,
		InOperation: c.PostForm("in_operation"),
		TransferFee: c.PostForm("transfer_fee"),
		Description: c.PostForm("description"),
		IndustryID:  industryID,
		QuotedPrice: c.PostForm("quoted_price"),
		ExplicitTel: explicit_tel,
		Remake:      c.PostForm("remake"),
		ShopName:    c.PostForm("shop_name"),
		Tel1:        c.PostForm("tel1"),
		Tel2:        c.PostForm("tel2"),
		Protect:     protect,
		BusType:     c.PostForm("bus_type"),
		ModelType:   c.PostForm("model_type"),
		MinArea:     min_area,
		MaxArea:     max_area,
		MinRent:     min_rent,
		MaxRent:     max_rent,
		Lots:        c.PostForm("lots"),
		SourceInfo:  c.PostForm("source_info"),
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res = adminService.EditQZQGProInfo(id, req)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加收款（物业列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param body body entity.PayInfoRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfo/payInfo [POST]
// @Security ApiKeyAuth
func AddPayInfoByProInfo(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayInfoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddPayInfo(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改收款（物业列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "物业ID"
// @Param body body entity.PayInfoRequestByProInfo true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfo/payInfo/{id} [PUT]
// @Security ApiKeyAuth
func EditPayInfoByProInfo(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayInfoRequestByProInfo{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditPayInfoByProInfo(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 收款详情（物业列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "物业ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfo/payInfo/{id} [GET]
// @Security ApiKeyAuth
func QueryPayInfoByProInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryPayInfoByProInfo(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改收款（收款列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "收款ID"
// @Param body body entity.PayInfoRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payInfo/{id} [PUT]
// @Security ApiKeyAuth
func EditPayInfo(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayInfoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditPayInfo(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 添加收款（收款列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param body body entity.PayInfoRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payInfo [GET]
// @Security ApiKeyAuth
func AddPayInfo(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayInfoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddPayInfo(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 收款详情（收款列表）
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payInfo/{id} [GET]
// @Security ApiKeyAuth
func QueryPayInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryPayInfo(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 收款列表
// @tags 后台
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param payee_id query string false "业绩归属人ID"
// @Param pay_methond_id query string false "支付方式ID"
// @Param pay_month query string false "按月查询"
// @Param pay_year query string false "按年查询"
// @Param name query string false "收款人"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payInfo [GET]
// @Security ApiKeyAuth
func QueryPayInfos(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	args := map[string]interface{}{
		"payee_id":       c.Query("payee_id"),
		"pay_methond_id": c.Query("pay_methond_id"),
		"pay_month":      c.Query("pay_month"),
		"pay_year":       c.Query("pay_year"),
		"name":           c.Query("name"),
	}
	res := adminService.QueryPayInfos(pageSize, page, args)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加收款方式
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param body body entity.PayMethondRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payMethond [POST]
// @Security ApiKeyAuth
func AddPayMethond(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayMethondRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		res = adminService.AddPayMethond(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改收款方式
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "收款方式ID"
// @Param body body entity.PayMethondRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payMethond/{id} [PUT]
// @Security ApiKeyAuth
func EditPayMethond(c *gin.Context) {
	res := entity.ResponseData{}
	req := entity.PayMethondRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = "请求参数JSON错误"
	} else {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		res = adminService.EditPayMethond(id, req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除支付
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payInfo/{ids} [DELETE]
// @Security ApiKeyAuth
func DelPayInfo(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelPayInfo(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 删除支付方式
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param ids path string true "ID 多个用,分开"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payMethond/{ids} [DELETE]
// @Security ApiKeyAuth
func DelPayMethond(c *gin.Context) {
	id := c.Param("ids")
	idArr := strings.Split(id, ",")
	var ids []int64
	for _, v := range idArr {
		item, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, item)
	}
	res := adminService.DelPayMethond(ids)
	c.JSON(http.StatusOK, res)
}

// @Summary 收款方式详情
// @tags 后台
// @Accept application/json
// @Produce  json
// @Param id path string false "ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payMethond/{id} [GET]
// @Security ApiKeyAuth
func QueryPayMethondByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.QueryPayMethondByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 收款方式列表
// @tags 后台
// @Accept application/json
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/payMethond [GET]
// @Security ApiKeyAuth
func QueryPayMethond(c *gin.Context) {
	res := adminService.QueryPayMethond()
	c.JSON(http.StatusOK, res)
}

// @Summary 获取用户菜单
// @tags 后台
// @Accept application/json
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/rolemenus [GET]
// @Security ApiKeyAuth
func QueryUserMenu(c *gin.Context) {
	var token string
	res := entity.ResponseData{}
	if token, res = commonController.GetToken(c); res.Status {
		res = adminService.QueryUserMenu(token)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 物业成功
// @tags 后台
// @Accept  application/json
// @Produce  json
// @Param id path string true "物业ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/admin/proInfos/success/{id} [PUT]
// @Security ApiKeyAuth
func EditProInfoSuccess(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := adminService.EditProInfoSuccess(id)
	c.JSON(http.StatusOK, res)
}
