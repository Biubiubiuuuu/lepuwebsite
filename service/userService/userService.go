package userService

import (
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/encryptHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/jwtHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/utilsHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/commonService"
	"github.com/google/uuid"
)

// 用户注册
func Register(req entity.UserRegister) (res entity.ResponseData) {
	if req.Telephone == "" || req.Password == "" {
		res.Message = "手机号码或密码不能为空"
		return
	}
	if req.Code == "" {
		res.Message = "验证码不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "手机号码格式不正确"
		return
	}
	if !utilsHelper.CheckPasswordFormat(req.Password) {
		res.Message = "密码格式不正确，密码可包含数字、英文、!@#$&*.,字符，长度6-20"
		return
	}
	/* 	v := model.Verificationcode{Tel: req.Telephone}
	   	if err := v.GetVerificationcode(); err != nil {
	   		res.Message = "验证码获取失败"
	   		return
	   	}
	   	t1 := utilsHelper.TimestampToTime(v.CreateTime)
	   	t2 := time.Now()
	   	sub := t2.Sub(t1)
	   	if sub.Seconds() > 60 {
	   		res.Message = "验证码已过期，请重新获取"
	   		return
	   	}
	   	if v.Code != req.Code {
	   		res.Message = "验证码错误"
	   		return
	   	} */
	uuid, _ := uuid.NewUUID()
	u := model.User{
		Username:  req.Telephone,
		Telephone: req.Telephone,
		Password:  encryptHelper.EncryptMD5To32Bit(req.Password),
		Sex:       "0",
		UUID:      uuid,
		Type:      "0",
	}
	if err := u.QueryByUsernameOrPhone(); err == nil {
		res.Message = "手机号码已注册，请直接登录"
		return
	}
	if err := u.Register(); err != nil {
		res.Message = "注册失败"
		return
	}
	res.Status = true
	res.Message = "注册成功"
	return
}

// 登录
func Login(req entity.UserLogin, ip string) (res entity.ResponseData) {
	if req.UserName == "" || req.Password == "" {
		res.Message = "用户名或密码不能为空"
		return
	}
	u := model.User{
		Username:  req.UserName,
		Telephone: req.UserName,
	}
	if err := u.QueryByUsernameOrPhone(); err != nil {
		res.Message = "用户名未注册，请先注册"
		return
	}
	pass := encryptHelper.EncryptMD5To32Bit(req.Password)
	if u.Password != pass {
		res.Message = "用户名或密码错误！"
		return
	}
	token, err := jwtHelper.GenerateToken(req.UserName, pass)
	if err != nil {
		res.Message = "登录失败，token生成错误！"
		return
	}
	args := map[string]interface{}{
		"token": token,
		"ip":    ip,
	}
	if err := u.Edit(args); err != nil {
		res.Message = "登录失败，更新登录信息失败！"
		return
	}
	data := make(map[string]interface{})
	data["user"] = u
	res.Status = true
	res.Message = "登录成功"
	res.Data = data
	return
}

// 修改用户信息
func EditUser(token string, req entity.EditUser) (res entity.ResponseData) {
	if strings.Trim(req.Username, " ") == "" {
		res.Message = "用户名不能为空！"
		return
	}
	if strings.Trim(req.Telephone, " ") == "" {
		res.Message = "手机号码不能为空！"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "手机号码格式不正确！"
		return
	}
	var user model.User
	if user, res = commonService.QueryUserByToken(token); !res.Status {
		return
	}

	if !(req.Sex == "0" || req.Sex == "1" || req.Sex == "3") {
		req.Sex = "0"
	}
	args := map[string]interface{}{
		"username":       req.Username,
		"telephone":      req.Telephone,
		"sex":            req.Sex,
		"email":          req.Email,
		"landlinenumber": req.Landlinenumber,
		"nickname":       req.Nickname,
		"QQ":             req.QQ,
	}
	u := model.User{
		Username: req.Username,
	}
	if err := u.QueryByUsername(); err == nil && u.Username != user.Username {
		res.Message = "修改失败，该用户名已注册"
		return
	}
	u = model.User{
		Telephone: req.Telephone,
	}
	if err := u.QueryByPhone(); err == nil && u.Telephone != user.Telephone {
		res.Message = "修改失败，该手机号已注册"
		return
	}
	if err := user.Edit(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 修改用户密码
func EditUserPass(token string, req entity.EditUserPass) (res entity.ResponseData) {
	if strings.Trim(req.NewPass, " ") == "" || strings.Trim(req.OldPass, " ") == "" {
		res.Message = "密码不能为空！"
		return
	}
	if !utilsHelper.CheckPasswordFormat(req.NewPass) {
		res.Message = "密码格式不正确，6-12位，至少包含数字跟字母，可以有字符"
		return
	}
	var user model.User
	if user, res = commonService.QueryUserByToken(token); !res.Status {
		return
	}
	if user.Password != encryptHelper.EncryptMD5To32Bit(req.OldPass) {
		res.Message = "旧密码错误"
		return
	}
	args := map[string]interface{}{
		"password": encryptHelper.EncryptMD5To32Bit(req.NewPass),
	}
	if err := user.Edit(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 查看用户信息
func QueryUserByToken(token string) (res entity.ResponseData) {
	var user model.User
	if user, res = commonService.QueryUserByToken(token); !res.Status {
		return
	}
	data := make(map[string]interface{})
	data["user"] = user
	res.Status = true
	res.Message = "获取成功"
	res.Data = data
	return
}

// 用户店铺转让
func UserStoretransfer(token string, req entity.UserStoretransferRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	if pros := store.QueryPropertyInfoByUserID(); len(pros) > 0 {
		res.Message = "你已经发布物业信息，请不要重复提交"
		return
	}
	industry := model.Industry{}
	industry.ID = req.IndustryID
	if err := industry.QueryIndustryByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	if req.Title == "" {
		res.Message = "标题不能为空"
		return
	}
	if req.Area <= 0 {
		res.Message = "面积必须大于0"
		return
	}
	if req.Rent <= 0 {
		res.Message = "租金必须大于0"
		return
	}
	province := model.Province{
		Code: req.ProvinceCode,
	}
	if err := province.QueryProvinceByCode(); err != nil {
		res.Message = "省不存在"
		return
	}
	city := model.City{
		Code:         req.CityCode,
		ProvinceCode: req.ProvinceCode,
	}
	if err := city.QueryCitysByCodeAndPro(); err != nil {
		res.Message = "该省份下城市不存在"
		return
	}
	district := model.District{
		Code:     req.DistrictCode,
		CityCode: req.CityCode,
	}
	if err := district.QueryDistrictByCodeAndCity(); err != nil {
		res.Message = "该城市下区不存在"
		return
	}
	street := model.Street{
		Code:         req.StreetCode,
		DistrictCode: req.Description,
	}
	if err := street.QueryStreetByCodeAndDist(); err != nil {
		street.Code = req.Description
		if arr := street.QueryStreetByDistrictCode(); len(arr) > 0 {
			res.Message = "该区下街道不存在"
			return
		}
	}
	if strings.Trim(req.Address, " ") == "" {
		res.Message = "详细地址不能为空"
		return
	}
	if strings.Trim(req.Telephone, " ") == "" {
		res.Message = "联系手机不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "联系手机格式不正确"
		return
	}
	if len(req.IndustryRanges) == 0 {
		res.Message = "适合经营至少选勾选一项"
		return
	}
	var industryRangeArr []model.IndustryRange
	for _, item := range req.IndustryRanges {
		ind := model.Industry{}
		ind.ID = item
		if err := ind.QueryIndustryByID(); err != nil {
			res.Message = "经营业态不存在"
			return
		}
		industryRange := model.IndustryRange{
			IndustryID:   item,
			IndustryName: ind.Name,
		}
		industryRangeArr = append(industryRangeArr, industryRange)
	}
	storeType := model.StoreType{}
	storeType.ID = req.StoreTypeID
	if err := storeType.QueryStoreTypeByID(); err != nil {
		res.Message = "店铺类型不存在"
		return
	}
	rent := model.RentType{}
	if err := rent.QueryRentTypeByRent(req.Rent); err != nil {
		res.Message = "租金类型不存在"
		return
	}
	area := model.AreaType{}
	if err := area.QueryAreaTypeByArea(req.Area); err != nil {
		res.Message = "面积类型不存在"
		return
	}
	pro := model.PropertyInfo{
		Title:          req.Title,
		IndustryID:     req.IndustryID,
		Description:    req.Description,
		Telephone:      req.Telephone,
		Nickname:       req.Nickname,
		Image:          req.Image,
		ProvinceCode:   req.ProvinceCode,
		CityCode:       req.CityCode,
		DistrictCode:   req.DistrictCode,
		StreetCode:     req.StreetCode,
		Address:        req.Address,
		RentTypeID:     rent.ID,
		AreaTypeID:     area.ID,
		Area:           req.Area,
		Rent:           req.Rent,
		Idling:         req.Idling,
		TransferFee:    req.TransferFee,
		ModelType:      "0",
		BusType:        "0",
		SourceID:       user.ID,
		IndustryRanges: industryRangeArr,
		StoreTypeID:    req.StoreTypeID,
	}
	if err := pro.CreatePropertyInfo(); err != nil {
		res.Message = "发布失败"
		return
	}
	res.Status = true
	res.Message = "发布成功，待管理员审核"
	return
}

// 用户我要找铺
func FindStore(token string, req entity.UserFindStoreRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	if pros := store.QueryPropertyInfoByUserID(); len(pros) > 0 {
		res.Message = "你已经发布物业信息，请不要重复提交"
		return
	}
	industry := model.Industry{}
	industry.ID = req.IndustryID
	if err := industry.QueryIndustryByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	storeType := model.StoreType{}
	storeType.ID = req.StoreTypeID
	if err := storeType.QueryStoreTypeByID(); err != nil {
		res.Message = "店铺类型不存在"
		return
	}
	var lots []model.Lot
	for _, item := range req.Lots {
		dis := model.District{
			Code: item,
		}
		if err := dis.QueryDistrictByCode(); err != nil {
			res.Message = "区代码不存在"
			return
		}
		lot := model.Lot{
			DistrictCode: dis.Code,
			DistrictName: dis.Name,
		}
		lots = append(lots, lot)
	}
	if req.Title == "" {
		res.Message = "标题不能为空"
		return
	}
	if req.Telephone == "" {
		res.Message = "联系手机不能为空"
		return
	}
	if req.MinRent < 0 || req.MaxArea < 0 || req.MaxRent < 0 || req.MinArea < 0 {
		res.Message = "租金或面积不能小于0"
		return
	}
	if req.MinRent > req.MaxRent {
		res.Message = "最小租金不能大于最大租金"
		return
	}
	if req.MinArea > req.MaxArea {
		res.Message = "最小面积不能大于最大面积"
		return
	}
	pro := model.PropertyInfo{
		StoreTypeID: req.StoreTypeID,
		IndustryID:  req.IndustryID,
		Lots:        lots,
		MinArea:     req.MinArea,
		MinRent:     req.MinRent,
		MaxArea:     req.MaxArea,
		MaxRent:     req.MaxRent,
		Title:       req.Title,
		Nickname:    req.Nickname,
		Telephone:   req.Telephone,
		Description: req.Description,
		ModelType:   "4",
		BusType:     "0",
	}
	if err := pro.CreatePropertyInfo(); err != nil {
		res.Message = "发布失败"
		return
	}
	res.Status = true
	res.Message = "发布成功，待管理员审核"
	return
}

// 查询用户已发布物业信息
func QueryUserPropertyInfo(token string) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	data := make(map[string]interface{})
	pros := store.QueryPropertyInfoByUserID()
	if len(pros) == 0 {
		res.Message = "用户未发布物业信息"
		return
	}
	data["propertyInfo"] = pros[0]
	res.Status = true
	res.Data = data
	res.Message = "获取成功"
	return
}

// 用户修改店铺转让信息
func EditUserStoretransfer(token string, req entity.UserStoretransferRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	if pros := store.QueryPropertyInfoByUserID(); len(pros) == 0 {
		res.Message = "暂未发布物业信息，请先提交在修改"
		return
	} else {
		store.ID = pros[0].ID
	}
	industry := model.Industry{}
	industry.ID = req.IndustryID
	if err := industry.QueryIndustryByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	if req.Title == "" {
		res.Message = "标题不能为空"
		return
	}
	if req.Area <= 0 {
		res.Message = "面积必须大于0"
		return
	}
	if req.Rent <= 0 {
		res.Message = "租金必须大于0"
		return
	}
	province := model.Province{
		Code: req.ProvinceCode,
	}
	if err := province.QueryProvinceByCode(); err != nil {
		res.Message = "省不存在"
		return
	}
	city := model.City{
		Code:         req.CityCode,
		ProvinceCode: req.ProvinceCode,
	}
	if err := city.QueryCitysByCodeAndPro(); err != nil {
		res.Message = "该省份下城市不存在"
		return
	}
	district := model.District{
		Code:     req.DistrictCode,
		CityCode: req.CityCode,
	}
	if err := district.QueryDistrictByCodeAndCity(); err != nil {
		res.Message = "该城市下区不存在"
		return
	}
	street := model.Street{
		Code:         req.StreetCode,
		DistrictCode: req.Description,
	}
	if err := street.QueryStreetByCodeAndDist(); err != nil {
		street.Code = req.Description
		if arr := street.QueryStreetByDistrictCode(); len(arr) > 0 {
			res.Message = "该区下街道不存在"
			return
		}
	}
	if strings.Trim(req.Address, " ") == "" {
		res.Message = "详细地址不能为空"
		return
	}
	if strings.Trim(req.Telephone, " ") == "" {
		res.Message = "联系手机不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "联系手机格式不正确"
		return
	}
	if len(req.IndustryRanges) == 0 {
		res.Message = "适合经营至少选勾选一项"
		return
	}
	var industryRangeArr []model.IndustryRange
	for _, item := range req.IndustryRanges {
		ind := model.Industry{}
		ind.ID = item
		if err := ind.QueryIndustryByID(); err != nil {
			res.Message = "经营业态不存在"
			return
		}
		industryRange := model.IndustryRange{
			IndustryID:   item,
			IndustryName: ind.Name,
		}
		industryRangeArr = append(industryRangeArr, industryRange)
	}
	storeType := model.StoreType{}
	storeType.ID = req.StoreTypeID
	if err := storeType.QueryStoreTypeByID(); err != nil {
		res.Message = "店铺类型不存在"
		return
	}
	rent := model.RentType{}
	if err := rent.QueryRentTypeByRent(req.Rent); err != nil {
		res.Message = "租金类型不存在"
		return
	}
	area := model.AreaType{}
	if err := area.QueryAreaTypeByArea(req.Area); err != nil {
		res.Message = "面积类型不存在"
		return
	}
	if !(req.InOperation == "0" || req.InOperation == "1" || req.InOperation == "2") {
		req.InOperation = "2"
	}
	args := map[string]interface{}{
		"title":           req.Title,
		"industry_id":     req.IndustryID,
		"description":     req.Description,
		"telephone":       req.Telephone,
		"nickname":        req.Nickname,
		"image":           req.Image,
		"province_code":   req.ProvinceCode,
		"city_code":       req.CityCode,
		"district_code":   req.DistrictCode,
		"street_code":     req.StreetCode,
		"address":         req.Address,
		"rent_type_id":    rent.ID,
		"area_type_id":    area.ID,
		"area":            req.Area,
		"rent":            req.Rent,
		"idling":          req.Idling,
		"transfer_fee":    req.TransferFee,
		"source_id":       user.ID,
		"industry_ranges": industryRangeArr,
		"store_type_id":   req.StoreTypeID,
	}
	if err := store.EditPropertyInfoByID(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 用户修改我要找铺信息
func EditUserFindStore(token string, req entity.UserFindStoreRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	if pros := store.QueryPropertyInfoByUserID(); len(pros) == 0 {
		res.Message = "暂未发布物业信息，请先提交在修改"
		return
	}
	industry := model.Industry{}
	industry.ID = req.IndustryID
	if err := industry.QueryIndustryByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	storeType := model.StoreType{}
	storeType.ID = req.StoreTypeID
	if err := storeType.QueryStoreTypeByID(); err != nil {
		res.Message = "店铺类型不存在"
		return
	}
	var lots []model.Lot
	for _, item := range req.Lots {
		dis := model.District{
			Code: item,
		}
		if err := dis.QueryDistrictByCode(); err != nil {
			res.Message = "区代码不存在"
			return
		}
		lot := model.Lot{
			DistrictCode: dis.Code,
			DistrictName: dis.Name,
		}
		lots = append(lots, lot)
	}
	if req.Title == "" {
		res.Message = "标题不能为空"
		return
	}
	if req.Telephone == "" {
		res.Message = "联系手机不能为空"
		return
	}
	if req.MinRent < 0 || req.MaxArea < 0 || req.MaxRent < 0 || req.MinArea < 0 {
		res.Message = "租金或面积不能小于0"
		return
	}
	if req.MinRent > req.MaxRent {
		res.Message = "最小租金不能大于最大租金"
		return
	}
	if req.MinArea > req.MaxArea {
		res.Message = "最小面积不能大于最大面积"
		return
	}
	args := map[string]interface{}{
		"store_type_id": req.StoreTypeID,
		"industry_id":   req.IndustryID,
		"lots":          lots,
		"min_area":      req.MinArea,
		"min_rent":      req.MinRent,
		"max_area":      req.MaxArea,
		"max_rent":      req.MaxRent,
		"title":         req.Title,
		"nickname":      req.Nickname,
		"telephone":     req.Telephone,
		"description":   req.Description,
	}
	if err := store.EditPropertyInfoByID(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 查询物业详情
func QueryPropertyInfoByID(id int64) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	data := map[string]interface{}{
		"propertyInfo": pro,
	}
	res.Status = true
	res.Data = data
	res.Message = "查询成功"
	return
}

// 精准匹配
func SearchPropertyInfo(pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
	pros, count := model.QueryPropertyInfo(pageSize, page, args)
	data := map[string]interface{}{
		"propertyInfos": pros,
		"count":         count,
	}
	res.Data = data
	res.Message = "查询成功"
	res.Status = true
	return
}

// 用户店铺转让上传图集
func AddPictures(token string, id int64, url string) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	is := false
	if pros := store.QueryPropertyInfoByUserID(); len(pros) == 0 {
		res.Message = "暂未发布物业信息，请先提交在修改"
		return
	} else {
		for _, item := range pros {
			if item.ID == id {
				is = true
				break
			}
		}
	}
	if !is {
		res.Message = "不存在该物业信息"
		return
	}
	picture := model.Picture{
		PropertyInfoID: id,
		Url:            url,
	}
	if err := picture.AddPicture(); err != nil {
		res.Message = "图片上传失败"
		return
	}
	res.Status = true
	res.Message = "上传成功"
	return
}

// 物业图集图片删除
func DelPrictures(token string, pro_id int64, pri_id int64) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	store := model.PropertyInfo{
		SourceID: user.ID,
	}
	is := false
	if pros := store.QueryPropertyInfoByUserID(); len(pros) == 0 {
		res.Message = "暂未发布物业信息，请先提交在修改"
		return
	} else {
		for _, item := range pros {
			if item.ID == pro_id {
				is = true
				break
			}
		}
	}
	if !is {
		res.Message = "不存在该物业信息"
		return
	}
	pri := model.Picture{
		PropertyInfoID: pro_id,
	}
	pri.ID = pri_id
	if err := pri.DelPicturre(); err != nil {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = "删除成功"
	return
}
