package main

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/configHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/encryptHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
	"github.com/Biubiubiuuuu/yuepuwebsite/router"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//初始化mysql
	mysql.DB.InitConn()
	db := mysql.GetMysqlDB()
	//自动迁移模型
	db.AutoMigrate(
		&model.User{},
		&model.Industry{},
		&model.AreaType{},
		&model.RentType{},
		&model.Province{},
		&model.City{},
		&model.District{},
		&model.Street{},
		&model.PropertyInfo{},
		&model.LeaveMessage{},
		&model.ShopTransferLog{},
		&model.StoreType{},
		&model.IndustryRange{},
		&model.Picture{},
		&model.Lot{},
	)
	// 添加默认管理员 username:Admin,password:123456
	u := model.User{Username: "admin", Password: encryptHelper.EncryptMD5To32Bit("123456"), Type: "1"}
	if err := u.QueryByUsernameOrPhone(); err != nil {
		u.Register()
	}
	//初始化redis
	//redis.DB.InitConn()
	//初始化路由
	router := router.Init()
	//启动
	router.Run(configHelper.HTTPPort)
}
