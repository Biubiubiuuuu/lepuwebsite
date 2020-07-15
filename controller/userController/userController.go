package userController

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/configHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/fileHelper"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"

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