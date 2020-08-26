package router

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/controller/adminController"
	"github.com/Biubiubiuuuu/yuepuwebsite/controller/basicController"
	"github.com/Biubiubiuuuu/yuepuwebsite/controller/commonController"
	"github.com/Biubiubiuuuu/yuepuwebsite/controller/userController"
	"github.com/Biubiubiuuuu/yuepuwebsite/docs"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/configHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/middleware/adminMiddleware"
	"github.com/Biubiubiuuuu/yuepuwebsite/middleware/crossMiddleware"
	"github.com/Biubiubiuuuu/yuepuwebsite/middleware/errorMiddleware"
	"github.com/Biubiubiuuuu/yuepuwebsite/middleware/jwtMiddleware"
	"github.com/Biubiubiuuuu/yuepuwebsite/middleware/loggerMiddleware"
	"github.com/gin-gonic/gin"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化路由
func Init() *gin.Engine {
	// swagger接口文档
	docs.SwaggerInfo.Title = "乐铺网"
	docs.SwaggerInfo.Description = "乐铺网"
	docs.SwaggerInfo.Version = configHelper.Version
	// 设置模式
	if configHelper.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// 记录日志
	router.Use(loggerMiddleware.Logger())
	// 静态资源路径 /static 开头 或者 取自定义配置
	router.Static(configHelper.Static, "."+configHelper.Static)
	//router.Static("/static", "./static")
	//允许跨域请求
	router.Use(crossMiddleware.Cors())

	// 自定义router
	InitUser(router)
	InitCommon(router)
	InitBasic(router)
	InitAdmin(router)
	InitPropertyInfo(router)
	//gin swaager
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
	//404
	router.NoRoute(errorMiddleware.NotFound)
	return router
}

// 物业信息接口
func InitPropertyInfo(router *gin.Engine) {
	api := router.Group("api/v1/propertyInfo")
	api.GET("", userController.SearchPropertyInfo)
	api.GET(":id", userController.QueryPropertyInfoByID)
}

// 用户
func InitUser(router *gin.Engine) {
	api := router.Group("api/v1/user")
	api.POST("login", userController.Login)
	api.POST("register", userController.Register)
	api.Use(jwtMiddleware.JWT())
	{
		api.GET("", userController.QueryUser)
		api.POST("", userController.EditUser)
		api.POST("edituserpass", userController.EditUserPass)
		api.POST("storetransfer", userController.Storetransfer)
		api.PUT("storetransfer", userController.EditUserStoretransfer)
		api.POST("findStore", userController.FindStore)
		api.PUT("findStore", userController.EditUserFindStore)
		api.GET("propertyInfo", userController.QueryUserPropertyInfo)
		api.POST("propertyInfo/:id/picture", userController.AddPictures)
		api.DELETE("propertyInfo/:pro_id/picture/:pri_id", userController.DelPrictures)
	}
	api2 := router.Group("api/v1/other")
	api2.POST("leaveMessage", userController.AddLeaveMessage)
	api2.POST("propertyInfo/:pro_id/report", userController.AddReport)
	api2.GET("carouse", userController.QueryCarouse)
	api2.GET("advert", userController.QueryAdvert)
}

// 公共接口
func InitCommon(router *gin.Engine) {
	api := router.Group("/api/v1/common")
	api.GET("verificationcode", commonController.VerificationCode)
}

// 基础数据
func InitBasic(router *gin.Engine) {
	api := router.Group("/api/v1/basic")
	api.GET("province", basicController.QueryProvinces)
	api.GET("city", basicController.QueryCitysByProvinceCode)
	api.GET("district", basicController.QueryDistrictByCityCode)
	api.GET("districts", basicController.QueryDistrict)
	api.GET("street", basicController.QueryStreetByDistrictCode)
	api.GET("streets", basicController.QueryStreet)
	api.GET("areatype", basicController.QueryAreaType)
	api.GET("renttype", basicController.QueryRentType)
	api.GET("enableStoreType", basicController.QueryEnableStoreType)
	api.GET("enableIndustry", basicController.QueryEnableIndustry)
	api.GET("industryParent", basicController.QueryEnableIndustryByParentID)
	api.GET("industrys", basicController.QueryIndustryByParentID)
	api.GET("proInfoDynamic", basicController.QueryProInfoDynamic)
	api.GET("systemConfig", basicController.QuerySystemConfigByDefault)
}

// 后台
func InitAdmin(router *gin.Engine) {
	api := router.Group("api/v1/admin")
	api.POST("login", adminController.Login)
	api.Use(adminMiddleware.JWTAndAdmin())
	{
		api.GET("", adminController.QueryUser)
		api.POST("edituserpass", userController.EditUserPass)

		api.GET("industry/:id", adminController.QueryIndustryByID)
		api.GET("industry", adminController.QueryIndustry)
		api.POST("industry", adminController.AddIndustry)
		api.PUT("industry/:id", adminController.EditIndustry)
		api.DELETE("industry/:ids", adminController.DelIndustry)

		api.GET("storeType/:id", adminController.QueryStoreTypeByID)
		api.GET("storeType", adminController.QueryStoreType)
		api.POST("storeType", adminController.AddStoreType)
		api.PUT("storeType/:id", adminController.EditStoreType)
		api.DELETE("storeType/:ids", adminController.DelStoreType)

		api.GET("areaType/:id", adminController.QueryAreaTypeInfoById)
		api.GET("areaType", adminController.QueryAreaType)
		api.POST("areaType", adminController.CreateAreaType)
		api.PUT("areaType/:id", adminController.EditAreaType)
		api.DELETE("areaType/:ids", adminController.DelAreaType)

		api.GET("rentType/:id", adminController.QueryRentTypeInfoById)
		api.GET("rentType", adminController.QueryRentType)
		api.POST("rentType", adminController.CreateRentType)
		api.PUT("rentType/:id", adminController.EditRentType)
		api.DELETE("rentType/:ids", adminController.DelRentType)

		api.GET("department/:id", adminController.QueryDepartmentByID)
		api.GET("department", adminController.QueryDepartment)
		api.POST("department", adminController.AddDepartment)
		api.PUT("department/:id", adminController.EditDepartment)
		api.DELETE("department/:ids", adminController.DelDepartment)

		api.GET("post/:id", adminController.QueryPostByID)
		api.GET("post", adminController.QueryPost)
		api.POST("post", adminController.AddPost)
		api.PUT("post/:id", adminController.EditPost)
		api.DELETE("post/:ids", adminController.DelPost)

		api.GET("role/:id", adminController.QueryRoleByID)
		api.GET("role", adminController.QueryRole)
		api.POST("role", adminController.AddRole)
		api.PUT("role/:id", adminController.EditRole)
		api.DELETE("role/:ids", adminController.DelRole)

		api.GET("menu/:id", adminController.QueryMenuByID)
		api.GET("menu", adminController.QueryMenu)
		api.POST("menu", adminController.AddMenu)
		api.PUT("menu/:id", adminController.Editmenu)
		api.DELETE("menu/:ids", adminController.DelMenu)

		api.GET("employee/:id", adminController.QueryEmployeeById)
		api.GET("employee", adminController.GetEmployee)
		api.POST("employee", adminController.AddEmployee)
		api.PUT("employee/:id", adminController.EditDepartment)
		api.DELETE("employee/:ids", adminController.DelEmployee)

		api.GET("propertyInfo/:id", adminController.QueryPropertyInfoByID)
		api.GET("propertyInfo", adminController.SearchPropertyInfo)
		api.PUT("propertyInfo/:id", adminController.EditUserStoretransfer)
		api.POST("propertyInfo/protect/:id", adminController.EditProtectionProInfo)
		api.POST("propertyInfo/notprotect/:id", adminController.EditNotProtectionProInfo)
		api.POST("propertyInfos/:id/picture", adminController.AddPictures)
		api.DELETE("propertyInfos/:pro_id/picture/:pri_id", adminController.DelPrictures)
		api.PUT("proInfos/success/:id", adminController.EditProInfoSuccess)

		api.GET("proInfo/payInfo/:id", adminController.QueryPayInfoByProInfo)
		api.POST("proInfo/payInfo", adminController.AddPayInfoByProInfo)
		api.PUT("proInfo/payInfo/:id", adminController.EditPayInfoByProInfo)

		// 跟单记录
		api.GET("proInfo/log/:id", adminController.QueryProInfoLog)
		api.POST("proInfo/log/:id", adminController.AddProInfoLog)

		api.POST("new/propertyInfo", adminController.AddProInfo)
		api.PUT("new/propertyInfo/:id", adminController.EditProInfo)
		api.DELETE("new/propertyInfo/:id", adminController.DelProInfo)
		api.POST("new/qzqgPropertyInfo", adminController.AddQZQGProInfo)
		api.PUT("new/qzqgPropertyInfo", adminController.EditQZQGProInfo)

		api.GET("leaveMessage/:id", adminController.QueryLeaveMessageByID)
		api.GET("leaveMessage", adminController.QueryLeaveMessage)

		api.GET("report/:id", adminController.QueryReportByID)
		api.GET("report", adminController.QueryReport)

		api.GET("advert/:id", adminController.QueryAdvertByID)
		api.GET("advert", adminController.QueryAdvert)
		api.POST("advert", adminController.AddAdvert)
		api.PUT("advert/:id", adminController.EditAdvert)
		api.DELETE("advert/:ids", adminController.DelAdvert)

		api.GET("carousel/:id", adminController.QueryCarouselByID)
		api.GET("carousel", adminController.QueryCarousel)
		api.POST("carousel", adminController.AddCarousel)
		api.PUT("carousel/:id", adminController.EditCarousel)
		api.DELETE("carousel/:ids", adminController.DelCarousel)

		api.GET("payMethond/:id", adminController.QueryPayMethondByID)
		api.GET("payMethond", adminController.QueryPayMethond)
		api.POST("payMethond", adminController.AddPayMethond)
		api.PUT("payMethond/:id", adminController.EditPayMethond)
		api.DELETE("payMethond/:ids", adminController.DelPayMethond)

		api.GET("payInfo/:id", adminController.QueryPayInfo)
		api.GET("payInfo", adminController.QueryPayInfos)
		api.POST("payInfo", adminController.AddPayInfo)
		api.PUT("payInfo/:id", adminController.EditPayInfo)
		api.DELETE("payInfo/:ids", adminController.DelPayInfo)

		api.GET("rolemenus", adminController.QueryUserMenu)
	}
}
