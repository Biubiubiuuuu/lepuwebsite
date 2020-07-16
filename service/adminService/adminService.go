package adminService

import (
	"fmt"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/encryptHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/jwtHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
)

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
	if u.Type != "1" {
		res.Message = "没有权限访问请求资源"
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

// 添加面积分类
func CreateAreaType(req entity.AreaTypeRequest) (res entity.ResponseData) {
	if req.MinArea < 0 {
		res.Message = "最小面积不能为0"
		return
	}
	if req.MaxArea <= 0 {
		res.Message = "最大面积不能为0"
		return
	}
	if req.MinArea >= req.MaxArea {
		res.Message = "最小面积必须小于最大面积"
		return
	}
	areaType := model.AreaType{}
	if err := areaType.QueryMaxArea(); err == nil && areaType.MaxArea != req.MinArea {
		res.Message = fmt.Sprintf("最小面积必须等于已添加的最大面积：%v", areaType.MaxArea)
		return
	}
	a := model.AreaType{
		MaxArea: req.MaxArea,
		MinArea: req.MinArea,
	}
	if err := a.AddAreaType(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改面积分类
func EditAreaType(id int64, req entity.AreaTypeRequest) (res entity.ResponseData) {
	if req.MinArea < 0 {
		res.Message = "最小面积不能为0"
		return
	}
	if req.MaxArea < 0 {
		res.Message = "最大面积不能为0"
		return
	}
	if req.MinArea > req.MaxArea {
		res.Message = "最小面积不能大于最大面积"
		return
	}
	args := map[string]interface{}{
		"min_area": req.MinArea,
		"max_area": req.MaxArea,
	}
	a := model.AreaType{}
	a.ID = id
	if err := a.EditAreaType(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 删除面积分类
func DelAreaType(ids []int64) (res entity.ResponseData) {
	if len(ids) == 0 {
		res.Message = "ID不能为空"
		return
	}
	if model.QueryPropertyInfoRelationAreaType(ids) {
		res.Message = "物业信息已关联此面积分类，请直接修改面积分类信息"
		return
	}
	count := model.DelAreaType(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加租金分类
func CreateRentType(req entity.RentTypeRequest) (res entity.ResponseData) {
	if req.MinRent < 0 {
		res.Message = "最小租金不能小于0"
		return
	}
	if req.MaxRent <= 0 {
		res.Message = "最大租金不能小于0"
		return
	}
	if req.MinRent >= req.MaxRent {
		res.Message = "最小租金必须小于最大租金"
		return
	}
	rentType := model.RentType{}
	if err := rentType.QueryMaxRent(); err == nil && rentType.MaxRent != req.MinRent {
		res.Message = fmt.Sprintf("最小租金必须等于已添加的最大租金：%v", rentType.MaxRent)
		return
	}
	a := model.RentType{
		MaxRent: req.MaxRent,
		MinRent: req.MinRent,
	}
	if err := a.AddRentType(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改租金分类
func EditRentType(id int64, req entity.RentTypeRequest) (res entity.ResponseData) {
	if req.MinRent < 0 {
		res.Message = "最小租金不能小于0"
		return
	}
	if req.MaxRent < 0 {
		res.Message = "最大租金不能小于0"
		return
	}
	if req.MinRent > req.MaxRent {
		res.Message = "最小租金不能大于最大租金"
		return
	}
	args := map[string]interface{}{
		"min_area": req.MinRent,
		"max_area": req.MaxRent,
	}
	r := model.RentType{}
	r.ID = id
	if err := r.EditRentType(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 删除租金分类
func DelRentType(ids []int64) (res entity.ResponseData) {
	if len(ids) == 0 {
		res.Message = "ID不能为空"
		return
	}
	if model.QueryPropertyInfoRelationRentType(ids) {
		res.Message = "物业信息已关联此租金分类，请直接修改租金分类信息"
		return
	}
	count := model.DelRentType(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加经营业态
func AddIndustry(req entity.IndustryRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "经营业态名称不能为空"
		return
	}
	ind := model.Industry{
		Name: req.Name,
	}
	if err := ind.QueryIndustryByName(); err == nil {
		res.Message = "经营业态名称已存在"
		return
	}
	ind.ID = req.ParentID
	if err := ind.QueryIndustryByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级经营业态不存在"
		return
	}
	ind = model.Industry{
		Name:     req.Name,
		Sort:     req.Sort,
		IsEnable: req.IsEnable,
		ParentID: req.ParentID,
	}
	if err := ind.AddIndustry(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改经营业态
func EditIndustry(id int64, req entity.IndustryRequest) (res entity.ResponseData) {
	ind := model.Industry{}
	ind.ID = id
	if err := ind.QueryIndustryByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	if req.Name == "" {
		res.Message = "经营业态名称不能为空"
		return
	}
	ind2 := model.Industry{
		Name: req.Name,
	}
	if err := ind2.QueryIndustryByName(); err == nil && ind2.Name != ind.Name {
		res.Message = "经营业态名称已存在"
		return
	}
	ind3 := model.Industry{}
	ind3.ID = req.ParentID
	if err := ind3.QueryIndustryByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级经营业态不存在"
		return
	}
	args := map[string]interface{}{
		"name":      req.Name,
		"sort":      req.Sort,
		"is_enable": req.IsEnable,
		"parent_id": req.ParentID,
	}
	if err := ind.EditIndustry(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 删除经营业态
func DelIndustry(ids []int64) (res entity.ResponseData) {
	if model.QueryPropertyInfoRelationIndustry(ids) {
		res.Message = "物业信息已关联此经营业态，请直接修改经营业态信息"
		return
	}
	count := model.DelIndustry(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加店铺类型
func AddStoreType(req entity.StoreTypeRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "店铺类型名称不能为空"
		return
	}
	sto := model.StoreType{
		Name: req.Name,
	}
	if err := sto.QueryStoreTypeByName(); err == nil {
		res.Message = "店铺类型名称已存在"
		return
	}
	sto = model.StoreType{
		Name:     req.Name,
		Sort:     req.Sort,
		IsEnable: req.IsEnable,
	}
	if err := sto.AddStoreType(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改店铺类型
func EditStoreType(id int64, req entity.StoreTypeRequest) (res entity.ResponseData) {
	sto := model.StoreType{}
	sto.ID = id
	if err := sto.QueryStoreTypeByID(); err != nil {
		res.Message = "经营业态不存在"
		return
	}
	if req.Name == "" {
		res.Message = "经营业态名称不能为空"
		return
	}
	sto2 := model.StoreType{
		Name: req.Name,
	}
	if err := sto2.QueryStoreTypeByName(); err == nil && sto2.Name != sto.Name {
		res.Message = "经营业态名称已存在"
		return
	}
	args := map[string]interface{}{
		"name":      req.Name,
		"sort":      req.Sort,
		"is_enable": req.IsEnable,
	}
	if err := sto.EditStoreType(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 删除店铺类型
func DelStoreType(ids []int64) (res entity.ResponseData) {
	if model.QueryPropertyInfoRelationStoreTypeID(ids) {
		res.Message = "物业信息已关联此店铺类型，请直接修改店铺类型信息"
		return
	}
	count := model.DelStoreType(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}
