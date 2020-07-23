package userController

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/helper/configHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/fileHelper"
	"github.com/google/uuid"

	"github.com/Biubiubiuuuu/yuepuwebsite/controller/commonController"
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/userService"

	"github.com/gin-gonic/gin"
)

// @Summary 注册
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserRegister true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/register [POST]
func Register(c *gin.Context) {
	req := entity.UserRegister{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = userService.Register(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 用户登录
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserLogin true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/login [POST]
func Login(c *gin.Context) {
	req := entity.UserLogin{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = userService.Login(req, c.ClientIP())
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改用户信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.EditUser true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user [POST]
// @Security ApiKeyAuth
func EditUser(c *gin.Context) {
	req := entity.EditUser{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		var token string
		if token, res = commonController.GetToken(c); res.Status {
			res = userService.EditUser(token, req)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改用户密码
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.EditUserPass true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/edituserpass [POST]
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
// @tags 用户
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user [GET]
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

// @Summary 店铺转让
// @tags 用户
// @Accept  multipart/form-data
// @Produce  json
// @Param industry_id formData string true "经营业态ID"
// @Param store_type_id formData string true "店铺类型"
// @Param idling formData bool true "可否空转"
// @Param area formData string true "面积（单位：平方米）"
// @Param rent formData string true "租金（单位：元/月）"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param image formData file false "图片"
// @Param title formData string true "标题"
// @Param nickname formData string true "联系人"
// @Param telephone formData string true "联系手机"
// @Param province_code formData string true "省代码""
// @Param city_code formData string true "城市代码"
// @Param district_code formData string true "区代码"
// @Param street_code formData string false "街道代码"
// @Param address formData string true "详细地址"
// @Param description formData string false "详细描述""
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/storetransfer [POST]
// @Security ApiKeyAuth
func Storetransfer(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
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
			} else {
				res.Message = "图片格式只支持png、jpg"
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
		req := entity.UserStoretransferRequest{
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
		}
		res = userService.UserStoretransfer(token, req)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 我要找铺
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserFindStoreRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/findStore [POST]
// @Security ApiKeyAuth
func FindStore(c *gin.Context) {
	req := entity.UserFindStoreRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		var token string
		if token, res = commonController.GetToken(c); res.Status {
			res = userService.FindStore(token, req)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 查看用户已发布物业信息
// @tags 用户
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/propertyInfo [GET]
// @Security ApiKeyAuth
func QueryUserPropertyInfo(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
		res = userService.QueryUserPropertyInfo(token)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改店铺转让信息
// @tags 用户
// @Accept  multipart/form-data
// @Produce  json
// @Param industry_id formData string true "经营业态ID"
// @Param store_type_id formData string true "店铺类型"
// @Param idling formData bool true "可否空转"
// @Param area formData string true "面积（单位：平方米）"
// @Param in_operation formData string true "是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中"
// @Param rent formData string true "租金（单位：元/月）"
// @Param transfer_fee formData string false "转让费用（单位：万元 不录入则显示为面议）"
// @Param industry_ranges formData string true "适合经营范围id，多个用，拼接"
// @Param image formData file false "图片"
// @Param title formData string true "标题"
// @Param nickname formData string true "联系人"
// @Param telephone formData string true "联系手机"
// @Param province_code formData string true "省代码""
// @Param city_code formData string true "城市代码"
// @Param district_code formData string true "区代码"
// @Param street_code formData string false "街道代码"
// @Param address formData string true "详细地址"
// @Param description formData string false "详细描述""
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/storetransfer [PUT]
// @Security ApiKeyAuth
func EditUserStoretransfer(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
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
			} else {
				res.Message = "图片格式只支持png、jpg"
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
		req := entity.UserStoretransferRequest{
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
		}
		res = userService.EditUserStoretransfer(token, req)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 修改我要找铺信息
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.UserFindStoreRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/findStore [PUT]
// @Security ApiKeyAuth
func EditUserFindStore(c *gin.Context) {
	req := entity.UserFindStoreRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		var token string
		if token, res = commonController.GetToken(c); res.Status {
			res = userService.EditUserFindStore(token, req)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 根据条件获取物业信息
// @tags 物业信息
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
// @Param model_type query string true "模型类型 0-转让 ｜ 1-出售 ｜ 3-出租 | 4-求租 ｜ 5-求购 多个用，隔开"
// @Param bus_type query string false "业务类型 0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库"
// @Param sort_condition query string false "排序 area-面积 ｜ rent-租金 ｜ created_at-发布时间（默认）"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/propertyInfo [GET]
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
	res := userService.SearchPropertyInfo(pageSize, page, args)
	c.JSON(http.StatusOK, res)
}

// @Summary 查看物业详情
// @tags 物业信息
// @Accept  application/json
// @Produce  json
// @Param id path string true "物业信息ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/propertyInfo/{id} [GET]
func QueryPropertyInfoByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := userService.QueryPropertyInfoByID(id)
	c.JSON(http.StatusOK, res)
}

// @Summary 上传图集图片（单张）
// @tags 用户
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "物业信息ID"
// @Param image formData file false "图片"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/propertyInfo/{id}/picture [POST]
// @Security ApiKeyAuth
func AddPictures(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
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
			} else {
				res.Message = "图片格式只支持png、jpg、jpeg"
				c.JSON(http.StatusOK, res)
				return
			}
			fmt.Println(filename)
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
			res = userService.AddPictures(token, id, image)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 删除图片（单张）
// @tags 用户
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param pro_id path string true "物业信息ID"
// @Param pri_id path string true "图片ID"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/user/propertyInfo/{pro_id}/picture/{pri_id} [DELETE]
// @Security ApiKeyAuth
func DelPrictures(c *gin.Context) {
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
		pro_id, _ := strconv.ParseInt(c.Param("pro_id"), 10, 64)
		pri_id, _ := strconv.ParseInt(c.Param("pri_id"), 10, 64)
		res = userService.DelPrictures(token, pro_id, pri_id)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 留言
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param body body entity.LeaveMessageRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/other/leaveMessage [POST]
func AddLeaveMessage(c *gin.Context) {
	req := entity.LeaveMessageRequest{}
	res := entity.ResponseData{}
	if c.ShouldBindJSON(&req) != nil {
		res.Message = "请求参数json错误"
	} else {
		res = userService.AddLeaveMessage(req)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 举报
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param pro_id path string true "物业信息ID"
// @Param body body entity.ReportRequest true "body"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/other/propertyInfo/{pro_id}/report [POST]
// @Security ApiKeyAuth
func AddReport(c *gin.Context) {
	req := entity.ReportRequest{}
	res := entity.ResponseData{}
	var token string
	if token, res = commonController.GetToken(c); res.Status {
		if c.ShouldBindJSON(&req) != nil {
			res.Message = "请求参数json错误"
		} else {
			pro_id, _ := strconv.ParseInt(c.Param("pro_id"), 10, 64)
			res = userService.AddReport(token, pro_id, req)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 首页轮播
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/other/carouse [GET]
func QueryCarouse(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := userService.QueryCarouse(pageSize, page)
	c.JSON(http.StatusOK, res)
}

// @Summary 广告查询
// @tags 用户
// @Accept  application/json
// @Produce  json
// @Param hot query bool false "首页最热推广"
// @Param floor query bool false "F楼"
// @Param type query string false "信息列表推广 1-一栏四分之一图片广告 | 2-二栏四分之一图片广告 | 3-三栏重点推荐 | 4-五栏框架广告"
// @Param pageSize query string false "页大小 （默认30）"
// @Param page query string false "跳转页码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/other/advert [GET]
func QueryAdvert(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	res := userService.QueryAdvert(pageSize, page, map[string]interface{}{
		"hot":   c.Query("hot"),
		"floor": c.Query("floor"),
		"type":  c.Query("type"),
	})
	c.JSON(http.StatusOK, res)
}
