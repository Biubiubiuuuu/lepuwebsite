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
	//router.Use(loggerMiddleware.Logger())
	// 静态资源路径 /static 开头 或者 取自定义配置
	//router.Static(configHelper.Static, "." + configHelper.Static)
	router.Static("/static", "./static")
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
		api.POST("findStore", userController.FindStore)
		api.PUT("storetransfer", userController.EditUserStoretransfer)
		api.PUT("findStore", userController.EditUserFindStore)
		api.GET("propertyInfo", userController.QueryUserPropertyInfo)
	}
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
	api.GET("areatype", basicController.QueryAreaType)
	api.GET("renttype", basicController.QueryRentType)
	api.GET("enableStoreType", basicController.QueryEnableStoreType)
	api.GET("storeType", basicController.QueryStoreType)
	api.GET("enableIndustry", basicController.QueryEnableIndustry)
	api.GET("industry", basicController.QueryIndustry)
	api.GET("industryRange", basicController.QueryEnableIndustryByParentID)
}

// 后台
func InitAdmin(router *gin.Engine) {
	api := router.Group("api/v1/admin")
	api.POST("login", adminController.Login)
	api.Use(jwtMiddleware.JWT(), adminMiddleware.UserTypeIsAdmin())
	{
		api.POST("areaType", adminController.CreateAreaType)
		api.PUT("areaType/:id", adminController.EditAreaType)
		api.DELETE("areaType/:ids", adminController.DelAreaType)
		api.POST("rentType", adminController.CreateRentType)
		api.PUT("rentType/:id", adminController.EditRentType)
		api.DELETE("rentType/:ids", adminController.DelRentType)
		api.POST("industry", adminController.AddIndustry)
		api.PUT("industry/:id", adminController.EditIndustry)
		api.DELETE("industry/:ids", adminController.DelIndustry)
		api.POST("storeType", adminController.AddStoreType)
		api.PUT("storeType/:id", adminController.EditStoreType)
		api.DELETE("storeType/:ids", adminController.DelStoreType)
	}
}
