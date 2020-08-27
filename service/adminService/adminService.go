package adminService

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/encryptHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/jwtHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/utilsHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/commonService"
	"github.com/google/uuid"
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

// 查询面积分类详情
func QueryAreaTypeInfoById(id int64) (res entity.ResponseData) {
	areaType := model.AreaType{}
	areaType.ID = id
	if err := areaType.QueryAreaTypeByID(); err != nil {
		res.Message = "面积分类不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
	res.Data = map[string]interface{}{
		"areaType": areaType,
	}
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
	if err := a.QueryAreaTypeByID(); err != nil {
		res.Message = "面积分类不存在"
		return
	}
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

// 查询租金分类详情
func QueryRentTypeInfoById(id int64) (res entity.ResponseData) {
	rentType := model.RentType{}
	rentType.ID = id
	if err := rentType.QueryRentTypeInfoById(); err != nil {
		res.Message = "租金分类不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
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
		"min_rent": req.MinRent,
		"max_rent": req.MaxRent,
	}
	r := model.RentType{}
	r.ID = id
	if err := r.QueryRentTypeInfoById(); err != nil {
		res.Message = "租金分类不存在"
		return
	}
	if err := r.EditRentType(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 查询面积分类
func QueryAreaType() (res entity.ResponseData) {
	data := make(map[string]interface{})
	a := model.AreaType{}
	area_types := a.QueryAreaType()
	data["area_types"] = area_types
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询租金分类
func QueryRentType() (res entity.ResponseData) {
	data := make(map[string]interface{})
	r := model.RentType{}
	rent_types := r.QueryRentType()
	data["rent_types"] = rent_types
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
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

// 查询经营业态详情
func QueryIndustryByID(id int64) (res entity.ResponseData) {
	in := model.Industry{}
	in.ID = id
	if err := in.QueryIndustryByID(); err != nil {
		res.Message = "行业不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
	data := map[string]interface{}{
		"industry": in,
	}
	res.Data = data
	return
}

// 查询店铺类型详情
func QueryStoreTypeByID(id int64) (res entity.ResponseData) {
	sto := model.StoreType{}
	sto.ID = id
	if err := sto.QueryStoreTypeByID(); err != nil {
		res.Message = "行业不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
	data := map[string]interface{}{
		"storeType": sto,
	}
	res.Data = data
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

// 查询所有店铺类型
func QueryStoreType(pageSize int, page int, name string, enable string) (res entity.ResponseData) {
	data := make(map[string]interface{})
	count, store_types := model.QueryStoreType(pageSize, page, name, enable)
	data["store_types"] = store_types
	data["count"] = count
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 查询所有行业类型
func QueryIndustry(pageSize int, page int, name string, enable string) (res entity.ResponseData) {
	data := make(map[string]interface{})
	count, industrys := model.QueryIndustry(pageSize, page, name, enable)
	data["industrys"] = industrys
	data["count"] = count
	res.Data = data
	res.Status = true
	res.Message = "获取成功"
	return
}

// 添加部门
func AddDepartment(req entity.DepartmentRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "部门名称不能为空"
		return
	}
	dep := model.Department{}
	dep.ID = req.ParentID
	if err := dep.QueryDepartmentByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级部门不存在"
		return
	}
	department := model.Department{
		Name:     req.Name,
		ParentID: req.ParentID,
	}
	if err := department.QueryDepartmentByNameAndParentID(); err == nil {
		res.Message = "部门名称已存在"
		return
	}
	depart := model.Department{
		Name:     req.Name,
		ParentID: req.ParentID,
		Sort:     req.Sort,
		Enable:   req.Enable,
		Leading:  req.Leading,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	if err := depart.AddDepartment(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改部门信息
func EditDepartment(id int64, req entity.DepartmentRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "部门名称不能为空"
		return
	}
	dep := model.Department{}
	dep.ID = req.ParentID
	if err := dep.QueryDepartmentByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级部门不存在"
		return
	}
	department := model.Department{}
	department.ID = id
	if err := department.QueryDepartmentByID(); err != nil {
		res.Message = "修改部门不存在"
		return
	}
	args := map[string]interface{}{
		"name":      req.Name,
		"sort":      req.Sort,
		"leading":   req.Leading,
		"parent_id": req.ParentID,
		"email":     req.Email,
		"phone":     req.Phone,
		"enable":    req.Enable,
	}
	dep2 := model.Department{
		Name:     req.Name,
		ParentID: req.ParentID,
	}
	if req.ParentID == department.ID {
		args["parent_id"] = 0
	}
	if err := dep2.QueryDepartmentByNameAndParentID(); err != nil {
		if err2 := department.EditDepartmentByID(args); err2 != nil {
			res.Message = "修改失败"
			return
		}
		res.Message = "修改成功"
		res.Status = true
		return
	}
	if department.Name != req.Name {
		res.Message = "部门名称已存在"
		return
	}
	if err2 := department.EditDepartmentByID(args); err2 != nil {
		res.Message = "修改失败"
		return
	}
	res.Message = "修改成功"
	res.Status = true
	return
}

// 批量删除部门
func DelDepartment(ids []int64) (res entity.ResponseData) {
	if model.QueryUserByDepartmentID(ids) {
		res.Message = "用户已关联此部门信息，请直接修改"
		return
	}
	count := model.DelDepartments(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加岗位
func AddPost(req entity.PostRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "岗位名称不能为空"
		return
	}
	if req.Code == "" {
		res.Message = "岗位编码不能为空"
		return
	}
	post := model.Post{
		Code: req.Code,
	}
	if err := post.QueryPostByCode(); err == nil {
		res.Message = "岗位编码已存在"
		return
	}
	post = model.Post{
		Name: req.Name,
	}
	if err := post.QueryPostByName(); err == nil {
		res.Message = "岗位名称已存在"
		return
	}
	post = model.Post{
		Name:   req.Name,
		Code:   req.Code,
		Enable: req.Enable,
		Sort:   req.Sort,
	}
	if err := post.AddPost(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Message = "添加成功"
	res.Status = true
	return
}

// 修改岗位
func EditPost(id int64, req entity.PostRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "岗位名称不能为空"
		return
	}
	if req.Code == "" {
		res.Message = "岗位编码不能为空"
		return
	}
	post := model.Post{}
	post.ID = id
	if err := post.QueryPostByID(); err != nil {
		res.Message = "岗位信息不存在"
		return
	}
	post2 := model.Post{
		Code: req.Code,
		Name: req.Name,
	}
	if err := post2.QueryPostByCode(); err == nil && post2.Code != post.Code {
		res.Message = "岗位编码已存在"
		return
	}
	if err := post.QueryPostByName(); err == nil && post2.Name != post.Name {
		res.Message = "岗位名称已存在"
		return
	}
	args := map[string]interface{}{
		"name":   req.Name,
		"sort":   req.Sort,
		"code":   req.Code,
		"enable": req.Enable,
	}
	if err := post.EditPost(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 批量删除岗位
func DelPost(ids []int64) (res entity.ResponseData) {
	if model.QueryUserByPostID(ids) {
		res.Message = "用户已关联此岗位信息，请直接修改"
		return
	}
	count := model.DelPosts(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加角色
func AddRole(req entity.RoleRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "角色名称不能为空"
		return
	}
	role := model.Role{
		Name: req.Name,
	}
	if err := role.QueryRoleByName(); err == nil {
		res.Message = "角色名称已存在"
		return
	}
	var menuPowers []model.MenuPower
	for _, item := range req.MenuPowers {
		menu := model.Menu{}
		menu.ID = item.MenuID
		if err := menu.QueryMenuByID(); err != nil {
			res.Message = "菜单不存在"
			return
		}
		menuPower := model.MenuPower{
			MenuID:    item.MenuID,
			MenuTitle: menu.Title,
		}
		menuPowers = append(menuPowers, menuPower)
	}
	role = model.Role{
		Name:       req.Name,
		Sort:       req.Sort,
		Enable:     req.Enable,
		MenuPowers: menuPowers,
	}
	if err := role.AddRole(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Message = "添加成功"
	res.Status = true
	return
}

// 修改角色
func EditRole(id int64, req entity.RoleRequest) (res entity.ResponseData) {
	role := model.Role{}
	role.ID = id
	if err := role.QueryRoleByID(); err != nil {
		res.Message = "角色信息不存在"
		return
	}
	r := model.Role{
		Name: req.Name,
	}
	var menuPowers []model.MenuPower
	for _, item := range req.MenuPowers {
		menu := model.Menu{}
		menu.ID = item.MenuID
		if err := menu.QueryMenuByID(); err != nil {
			res.Message = "菜单不存在"
			return
		}
		menuPower := model.MenuPower{
			MenuID:    item.MenuID,
			MenuTitle: menu.Title,
		}
		menuPowers = append(menuPowers, menuPower)
	}
	args := map[string]interface{}{
		"name":        req.Name,
		"sort":        req.Sort,
		"enable":      req.Enable,
		"menu_powers": menuPowers,
	}
	if err := r.QueryRoleByName(); err != nil {
		if err := role.EditRole(args); err != nil {
			res.Message = "修改失败"
			return
		}
		res.Message = "修改成功"
		res.Status = true
		return
	}
	if r.ID != role.ID {
		res.Message = "角色名称已存在"
		return
	}
	if err := role.EditRole(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Message = "修改成功"
	res.Status = true
	return
}

// 批量删除角色
func DelRole(ids []int64) (res entity.ResponseData) {
	if model.QueryUserByRoleID(ids) {
		res.Message = "用户已关联此角色信息，请直接修改"
		return
	}
	count := model.DelRoles(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 查询部门详情
func QueryDepartmentByID(id int64) (res entity.ResponseData) {
	dep := model.Department{}
	dep.ID = id
	if err := dep.QueryDepartmentByID(); err != nil {
		res.Message = "部门不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
	data := map[string]interface{}{
		"department": dep,
	}
	res.Data = data
	return
}

// 查询岗位详细
func QueryPostByID(id int64) (res entity.ResponseData) {
	post := model.Post{}
	post.ID = id
	if err := post.QueryPostByID(); err != nil {
		res.Message = "岗位不存在"
		return
	}
	data := map[string]interface{}{
		"post": post,
	}
	res.Status = true
	res.Data = data
	res.Message = "查询成功"
	return
}

// 查询所有岗位
func QueryPost(pageSize int, page int, name string, code string, enable string) (res entity.ResponseData) {
	args := map[string]interface{}{
		"name":   name,
		"code":   code,
		"enable": enable,
	}
	count, posts := model.QueryPosts(pageSize, page, args)
	res.Status = true
	data := map[string]interface{}{
		"posts": posts,
		"count": count,
	}
	res.Data = data
	res.Message = "查询成功"
	return
}

// 查询部门
func QueryDepartment(pageSize int, page int, name string, enable string) (res entity.ResponseData) {
	count, dep := model.QueryDepartments(pageSize, page, name, enable)
	res.Status = true
	data := map[string]interface{}{
		"departments": dep,
		"count":       count,
	}
	res.Data = data
	res.Message = "查询成功"
	return
}

// 查询角色
func QueryRole(pageSize int, page int, name string, enable string) (res entity.ResponseData) {
	count, dep := model.QueryRole(pageSize, page, name, enable)
	res.Status = true
	data := map[string]interface{}{
		"roles": dep,
		"count": count,
	}
	res.Data = data
	res.Message = "查询成功"
	return
}

// 查询角色详细
func QueryRoleByID(id int64) (res entity.ResponseData) {
	role := model.Role{}
	role.ID = id
	if err := role.QueryRoleByID(); err != nil {
		res.Message = "角色不存在"
		return
	}
	data := map[string]interface{}{
		"role": role,
	}
	res.Status = true
	res.Data = data
	res.Message = "查询成功"
	return
}

// 添加菜单
func AddMenu(req entity.MenuRequest) (res entity.ResponseData) {
	if req.Title == "" {
		res.Message = "菜单名称不能为空"
		return
	}
	me := model.Menu{}
	me.ID = req.ParentID
	if err := me.QueryMenuByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级菜单不存在"
		return
	}
	menu := model.Menu{
		Title:      req.Title,
		Sort:       req.Sort,
		Icon:       req.Icon,
		ParentID:   req.ParentID,
		RouterName: req.RouterName,
		RouterUrl:  req.RouterUrl,
		Enable:     req.Enable,
	}
	if err := menu.AddMenu(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改菜单信息
func EditMenu(id int64, req entity.MenuRequest) (res entity.ResponseData) {
	if req.Title == "" {
		res.Message = "菜单名称不能为空"
		return
	}
	args := map[string]interface{}{
		"title":       req.Title,
		"sort":        req.Sort,
		"icon":        req.Icon,
		"parent_id":   req.ParentID,
		"router_name": req.RouterName,
		"router_url":  req.RouterUrl,
		"enable":      req.Enable,
	}
	menu := model.Menu{}
	menu.ID = req.ParentID
	if err := menu.QueryMenuByID(); err != nil && req.ParentID != 0 {
		res.Message = "上级菜单不存在"
		return
	}
	menu.ID = id
	if err := menu.QueryMenuByID(); err != nil {
		res.Message = "菜单不存在"
		return
	}
	if err := menu.EditMenu(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 批量删除菜单
func DelMenus(ids []int64) (res entity.ResponseData) {
	if ok := model.QueryRoleByMenuID(ids); ok {
		res.Message = "角色已添加菜单信息，请直接修改"
		return
	}
	count := model.DelMebus(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 查看菜单详情
func QueryMenuByID(id int64) (res entity.ResponseData) {
	menu := model.Menu{}
	menu.ID = id
	if err := menu.QueryMenuByID(); err != nil {
		res.Message = "菜单不存在"
		return
	}
	res.Status = true
	data := map[string]interface{}{
		"menu": menu,
	}
	res.Data = data
	res.Message = "查询成功"
	return
}

// 查询菜单
func QueryMenu() (res entity.ResponseData) {
	menus := model.QueryMenus()
	data := map[string]interface{}{
		"menus": menus,
	}
	res.Status = true
	res.Data = data
	res.Message = "查询成功"
	return
}

// 添加员工
func AddEmployee(req entity.AddEmployeeRequest) (res entity.ResponseData) {
	if req.Nickname == "" {
		res.Message = "用户昵称不能为空"
		return
	}
	if req.Phone == "" {
		res.Message = "用户手机号码不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Phone) {
		res.Message = "请输入有效的手机号码"
		return
	}
	if req.Password == "" {
		req.Password = "123456"
	}
	if !utilsHelper.CheckPasswordFormat(req.Password) {
		res.Message = "密码格式不正确，密码可包含数字、英文、!@#$&*.,字符，长度6-20"
		return
	}
	if req.Username == "" {
		res.Message = "用户名不能为空"
		return
	}
	if !(req.Sex == "0" || req.Sex == "1" || req.Sex == "3") {
		req.Sex = "0"
	}
	role := model.Role{}
	role.ID = req.RoleID
	if err := role.QueryRoleByID(); err != nil {
		res.Message = "角色不存在"
		return
	}
	dep := model.Department{}
	dep.ID = req.DepartmentID
	if err := dep.QueryDepartmentByID(); err != nil {
		res.Message = "部门不存在"
		return
	}
	post := model.Post{}
	post.ID = req.PostID
	if err := post.QueryPostByID(); err != nil {
		res.Message = "岗位不存在"
		return
	}
	if req.Email == "" {
		res.Message = "邮箱不能为空"
		return
	}
	user := model.User{
		Username: req.Username,
	}
	if err := user.QueryByUsername(); err == nil {
		res.Message = "用户名已存在"
		return
	}
	user.Telephone = req.Phone
	if err := user.QueryByPhone(); err == nil {
		res.Message = "手机号码已存在"
		return
	}
	uuid, _ := uuid.NewUUID()
	u := model.User{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Sex:       req.Sex,
		Password:  encryptHelper.EncryptMD5To32Bit(req.Password),
		Email:     req.Email,
		IsEnable:  req.Enable,
		Type:      "1",
		QQ:        req.QQ,
		Telephone: req.Phone,
		UUID:      uuid,
		UserInfo: model.UserInfo{
			DepartmentID:   req.DepartmentID,
			DepartmentName: dep.Name,
			RoleID:         req.RoleID,
			RoleName:       role.Name,
			PostID:         req.PostID,
			PostName:       post.Name,
		},
	}
	if err := u.Register(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
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

// 管理员修改用户信息
func EditAdminToUser(id int64, req entity.EditEmployeeRequest) (res entity.ResponseData) {
	user := model.User{}
	user.ID = id
	if err := user.QueryUserByID(); err != nil {
		res.Message = "用户信息不存在"
		return
	}
	if req.Nickname == "" {
		res.Message = "用户昵称不能为空"
		return
	}
	if req.Phone == "" {
		res.Message = "用户手机号码不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Phone) {
		res.Message = "请输入有效的手机号码"
		return
	}
	if req.Username == "" {
		res.Message = "用户名不能为空"
		return
	}
	if !(req.Sex == "0" || req.Sex == "1" || req.Sex == "3") {
		req.Sex = "0"
	}
	role := model.Role{}
	role.ID = req.RoleID
	if err := role.QueryRoleByID(); err != nil {
		res.Message = "角色不存在"
		return
	}
	dep := model.Department{}
	dep.ID = req.DepartmentID
	if err := dep.QueryDepartmentByID(); err != nil {
		res.Message = "部门不存在"
		return
	}
	post := model.Post{}
	post.ID = req.PostID
	if err := post.QueryPostByID(); err != nil {
		res.Message = "岗位不存在"
		return
	}
	if req.Email == "" {
		res.Message = "邮箱不能为空"
		return
	}
	user2 := model.User{
		Username: req.Username,
	}
	if err := user2.QueryByUsername(); err == nil && user.Username != req.Username {
		res.Message = "用户名已存在"
		return
	}
	user2.Telephone = req.Phone
	if err := user2.QueryByPhone(); err == nil && user.Telephone != req.Phone {
		res.Message = "手机号码已存在"
		return
	}
	args := map[string]interface{}{
		"username":  req.Username,
		"nickname":  req.Nickname,
		"sex":       req.Sex,
		"email":     req.Email,
		"is_enable": req.Enable,
		"user_info": model.UserInfo{
			DepartmentID:   req.DepartmentID,
			DepartmentName: dep.Name,
			RoleID:         req.RoleID,
			RoleName:       role.Name,
			PostID:         req.PostID,
			PostName:       post.Name,
		},
	}
	if err := user.Edit(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 重置密码
func NewAdminToPass(id int64) (res entity.ResponseData) {
	user := model.User{}
	user.ID = id
	if err := user.QueryUserByID(); err != nil {
		res.Message = "用户不存在"
		return
	}
	args := map[string]interface{}{
		"password": encryptHelper.EncryptMD5To32Bit("123456"),
	}
	if err := user.Edit(args); err != nil {
		res.Message = "重置失败"
		return
	}
	res.Status = true
	res.Message = "重置成功"
	return
}

// 查看用户详情 by id
func QueryEmployeeById(id int64) (res entity.ResponseData) {
	user := model.User{}
	user.ID = id
	if err := user.QueryEmployeeById(); err != nil {
		res.Message = "用户不存在"
		return
	}
	res.Status = true
	data := map[string]interface{}{
		"user": user,
	}
	res.Data = data
	res.Message = "查询成功"
	return
}

// 删除员工
func DelEmployee(ids []int64) (res entity.ResponseData) {
	count := model.DelEmployee(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 员工列表
func QueryAllUser(pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
	count, users := model.QueryUser(pageSize, page, args)
	if count == 0 {
		res.Message = "查询失败"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{
		"users": users,
		"count": count,
	}
	return
}

// 查看所有物业信息
func QueryAllPropertyInfo(pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
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

// 查询我的已发布信息
func SearchPropertyInfo(token string, pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "用户信息未找到"
		return
	}
	args["source_id"] = user.ID
	args["show_status"] = true
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

// 修改店铺转让信息
func EditUserStoretransfer(id int64, req entity.AdminStoretransferRequest) (res entity.ResponseData) {
	store := model.PropertyInfoScan{}
	store.ID = id
	if err := store.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	if req.Telephone == "" {
		res.Message = "联系手机号码不能为空"
		return
	}
	args := map[string]interface{}{
		"telephone": req.Telephone,
	}
	first, count := model.QueryPropertyInfo(1, 0, args)
	if count > 0 && first[0].Telephone != store.Telephone && first[0].ID != id {
		res.Message = "已存在该联系人物业信息"
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
	if !(req.InOperation == "0" || req.InOperation == "1" || req.InOperation == "2") {
		req.InOperation = "2"
	}
	args2 := map[string]interface{}{
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
		"industry_ranges": industryRangeArr,
		"store_type_id":   req.StoreTypeID,
		"explicit_tel":    req.ExplicitTel,
		"tel1":            req.Tel1,
		"tel2":            req.Tel2,
		"quoted_price":    req.QuotedPrice,
		"remake":          req.Remake,
	}
	if err := store.EditPropertyInfoByID(args2); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 保护
func EditProtectionProInfo(id int64) (res entity.ResponseData) {
	store := model.PropertyInfo{}
	store.ID = id
	args := map[string]interface{}{
		"protect": true,
	}
	if err := store.EditPropertyInfoByID(args); err != nil {
		res.Message = "保护失败"
		return
	}
	res.Status = true
	res.Message = "保护成功"
	return
}

// 取消保护
func EditNotProtectionProInfo(id int64) (res entity.ResponseData) {
	store := model.PropertyInfo{}
	store.ID = id
	args := map[string]interface{}{
		"protect": false,
	}
	if err := store.EditPropertyInfoByID(args); err != nil {
		res.Message = "取消保护失败"
		return
	}
	res.Status = true
	res.Message = "取消成功"
	return
}

// 上传图集
func AddPictures(id int64, url string) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
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
func DelPrictures(pro_id int64, pri_id int64) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = pro_id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
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

// 添加跟单记录
func AddProInfoLog(token string, id int64, req entity.AddProLog) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "用户不存在"
		return
	}
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	log := model.ProInfoLog{
		Content:   req.ContentText,
		WithID:    user.ID,
		WithName:  user.Nickname,
		ProInfoID: id,
	}
	if err := log.AddProInfoLog(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 查看跟单记录
func QueryByProInfoID(id int64) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	log := model.QueryByProInfoID(id)
	res.Message = "查询成功"
	res.Data = map[string]interface{}{"log": log}
	res.Status = true
	return
}

// 添加物业信息
func AddProInfo(token string, req entity.AddPropertyInfoRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "用户不存在"
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
	areaFlo, _ := strconv.ParseFloat(req.Area, 64)
	if areaFlo <= 0 {
		res.Message = "面积必须大于0"
		return
	}
	rentFlo, _ := strconv.ParseFloat(req.Rent, 64)
	if rentFlo <= 0 {
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
	_, count := model.QueryPropertyInfo(30, 1, map[string]interface{}{"telephone": req.Telephone})
	if count > 0 {
		res.Message = fmt.Sprintf("联系人手机号码%v已存在物业信息，请勿重复添加", req.Telephone)
		return
	}
	if req.IndustryRanges == "" {
		res.Message = "适合经营至少选勾选一项"
		return
	}
	var industryRangeArr []model.IndustryRange
	arr := strings.Split(req.IndustryRanges, ",")
	for _, item := range arr {
		ind := model.Industry{}
		id, _ := strconv.ParseInt(item, 10, 64)
		ind.ID = id
		if err := ind.QueryIndustryByID(); err != nil {
			res.Message = "经营业态不存在"
			return
		}
		industryRange := model.IndustryRange{
			IndustryID:   id,
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
	if err := rent.QueryRentTypeByRent(rentFlo); err != nil {
		res.Message = "租金类型不存在"
		return
	}
	area := model.AreaType{}
	if err := area.QueryAreaTypeByArea(areaFlo); err != nil {
		res.Message = "面积类型不存在"
		return
	}
	transferFee, _ := strconv.ParseFloat(req.TransferFee, 64)
	quotedPrice, _ := strconv.ParseFloat(req.QuotedPrice, 64)
	BusType, _ := strconv.ParseInt(req.BusType, 10, 64)
	ModelType, _ := strconv.ParseInt(req.ModelType, 10, 64)
	InOperation, _ := strconv.ParseInt(req.InOperation, 10, 64)
	pro := model.PropertyInfo{
		IndustryID:     req.IndustryID,
		Title:          req.Title,
		Nickname:       req.Nickname,
		Telephone:      req.Telephone,
		ShopName:       req.ShopName,
		Image:          req.Image,
		BusType:        BusType,
		ModelType:      ModelType,
		ProvinceCode:   req.ProvinceCode,
		CityCode:       req.CityCode,
		DistrictCode:   req.DistrictCode,
		StreetCode:     req.StreetCode,
		IndustryRanges: industryRangeArr,
		AreaTypeID:     area.ID,
		RentTypeID:     rent.ID,
		Area:           areaFlo,
		Rent:           rentFlo,
		Address:        req.Address,
		Description:    req.Description,
		Idling:         req.Idling,
		SourceID:       user.ID,
		Remake:         req.Remake,
		TransferFee:    transferFee,
		ExplicitTel:    req.ExplicitTel,
		Tel1:           req.Tel1,
		Tel2:           req.Tel2,
		InOperation:    InOperation,
		Protect:        req.Protect,
		StoreTypeID:    req.StoreTypeID,
		QuotedPrice:    quotedPrice,
	}
	if err := pro.CreatePropertyInfo(); err != nil {
		res.Message = "发布失败"
		return
	}
	res.Message = "发布成功"
	res.Status = true
	return
}

// 留言列表
func QueryLeaveMessage(pageSize int, page int) (res entity.ResponseData) {
	count, leaveMessages := model.QueryLeaveMessage(pageSize, page)
	if count == 0 {
		res.Message = "没有更多了"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{
		"leaveMessages": leaveMessages,
	}
	return
}

// 留言详情
func QueryLeaveMessageByID(id int64) (res entity.ResponseData) {
	mes := model.LeaveMessage{}
	mes.ID = id
	if err := mes.QueryLeaveMessageByID(); err != nil {
		res.Message = "留言不存在"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{
		"leaveMessage": mes,
	}
	return
}

// 举报列表
func QueryReport(pageSize int, page int) (res entity.ResponseData) {
	count, reports := model.QueryReport(pageSize, page)
	if count == 0 {
		res.Message = "没有更多了"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{
		"reports": reports,
	}
	return
}

// 举报详情
func QueryReportByID(id int64) (res entity.ResponseData) {
	report := model.Report{}
	report.ID = id
	if err := report.QueryReportByID(); err != nil {
		res.Message = "举报信息不存在"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{
		"report": report,
	}
	return
}

// 添加轮播图
func AddCarousel(req entity.CarouselRequest) (res entity.ResponseData) {
	carousel := model.Carousel{
		Url:  req.Url,
		Link: req.Link,
		Sort: req.Sort,
	}
	if err := carousel.AddCarousel(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改轮播
func EditCarousel(id int64, req entity.CarouselRequest) (res entity.ResponseData) {
	carousel := model.Carousel{}
	carousel.ID = id
	if err := carousel.QueryCarouselByID(); err != nil {
		res.Message = "未找到轮播信息"
		return
	}
	args := map[string]interface{}{
		"url":  req.Url,
		"linl": req.Link,
		"sort": req.Sort,
	}
	if err := carousel.EditCarousel(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 查询轮播详情
func QueryCarouselByID(id int64) (res entity.ResponseData) {
	carousel := model.Carousel{}
	carousel.ID = id
	if err := carousel.QueryCarouselByID(); err != nil {
		res.Message = "未找到轮播信息"
		return
	}
	res.Status = true
	res.Data = map[string]interface{}{"carousel": carousel}
	res.Message = "查询成功"
	return
}

// 查询轮播
func QueryCarousel(pageSize int, page int) (res entity.ResponseData) {
	count, carousels := model.QueryCarousel(pageSize, page)
	res.Status = true
	res.Data = map[string]interface{}{"carousels": carousels, "count": count}
	res.Message = "获取成功"
	return
}

// 删除轮播
func DelCarousel(ids []int64) (res entity.ResponseData) {
	count := model.DelCarousel(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加广告
func AddAdvert(req entity.AdvertRequest) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = req.PropertyInfoID
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	advertQue := model.Advert{
		PropertyInfoID: req.PropertyInfoID,
	}
	if err := advertQue.QueryAdvertByProID(); err == nil {
		res.Message = "此物业信息已添加广告，请直接修改"
		return
	}
	advert := model.Advert{
		PropertyInfoID: req.PropertyInfoID,
		StartTime:      utilsHelper.StringToSTime(req.StartTime),
		EndTime:        utilsHelper.StringToSTime(req.EndTime),
		Hot:            req.Hot,
		Floor:          req.Floor,
		Type:           req.Type,
		Sort:           req.Sort,
		Enable:         req.Enable,
	}
	if err := advert.AddAdvert(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Status = true
	res.Message = "添加成功"
	return
}

// 修改广告
func EditAdvert(id int64, req entity.AdvertRequest) (res entity.ResponseData) {
	advert := model.Advert{}
	advert.ID = id
	if err := advert.QueryAdvertByID(); err != nil {
		res.Message = "广告不存在"
		return
	}
	pro := model.PropertyInfoScan{}
	pro.ID = req.PropertyInfoID
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	advertQue := model.Advert{
		PropertyInfoID: req.PropertyInfoID,
	}
	if err := advertQue.QueryAdvertByProID(); err == nil && advertQue.ID != id {
		res.Message = "此物业信息已添加广告，请直接修改"
		return
	}
	args := map[string]interface{}{
		"sort":             req.Sort,
		"type":             req.Type,
		"floor":            req.Floor,
		"hot":              req.Hot,
		"end_time":         utilsHelper.StringToSTime(req.EndTime),
		"start_time":       utilsHelper.StringToSTime(req.StartTime),
		"property_info_id": req.PropertyInfoID,
		"enable":           req.Enable,
	}
	if err := advert.EditAdvertByID(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Message = "修改成功"
	res.Status = true
	return
}

// 查询广告
func QueryAdvert(pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
	count, adverts := model.QueryAdvert(pageSize, page, args)
	res.Status = true
	res.Data = map[string]interface{}{"adverts": adverts, "count": count}
	res.Message = "获取成功"
	return
}

// 查询广告详情
func QueryAdvertByID(id int64) (res entity.ResponseData) {
	advert := model.Advert{}
	advert.ID = id
	if err := advert.QueryAdvertByID(); err != nil {
		res.Message = "未找到广告信息"
		return
	}
	res.Status = true
	res.Data = map[string]interface{}{"advert": advert}
	res.Message = "查询成功"
	return
}

// 删除广告
func DelAdvert(ids []int64) (res entity.ResponseData) {
	count := model.DelAdvert(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 添加求租求购
func AddQZQGProInfo(token string, req entity.AddQZQGPropertyInfoRequest) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "用户不存在"
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
	if strings.Trim(req.Telephone, " ") == "" {
		res.Message = "联系手机不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "联系手机格式不正确"
		return
	}
	_, count := model.QueryPropertyInfo(30, 1, map[string]interface{}{"telephone": req.Telephone})
	if count > 0 {
		res.Message = fmt.Sprintf("联系人手机号码%v已存在物业信息，请勿重复添加", req.Telephone)
		return
	}
	city := model.City{
		Code: req.CityCode,
	}
	if err := city.QueryCitysByCode(); err != nil {
		res.Message = "城市代码不存在"
	}
	if !(req.BusType == "0" || req.BusType == "1" || req.BusType == "2") {
		req.BusType = "0"
	}
	if !(req.ModelType == "4" || req.ModelType == "5") {
		req.BusType = "4"
	}
	if req.MinArea <= 0 {
		res.Message = "最小面积必须大于0"
		return
	}
	if req.MaxArea <= 0 {
		res.Message = "最小面积必须大于0"
		return
	}
	if req.MinRent <= 0 {
		res.Message = "最小租金必须大于0"
		return
	}
	if req.MaxRent <= 0 {
		res.Message = "最小租金必须大于0"
		return
	}
	if req.MinArea > req.MaxArea {
		res.Message = "最小面积必须不大于最大面积"
		return
	}
	if req.MinRent > req.MaxRent {
		res.Message = "最小租金必须不大于最大租金"
		return
	}
	if req.Lots == "" {
		res.Message = "考虑地段至少选择一个"
		return
	}
	var lots []model.Lot
	arr := strings.Split(req.Lots, ",")
	for _, v := range arr {
		dis := model.District{
			Code: v,
		}
		if err := dis.QueryDistrictByCode(); err != nil {
			res.Message = "考虑地段区域不存在"
			return
		}
		lot := model.Lot{
			DistrictCode: dis.Code,
			DistrictName: dis.Name,
		}
		lots = append(lots, lot)
	}
	TransferFee, _ := strconv.ParseFloat(req.TransferFee, 64)
	QuotedPrice, _ := strconv.ParseFloat(req.QuotedPrice, 64)
	BusType, _ := strconv.ParseInt(req.BusType, 10, 64)
	ModelType, _ := strconv.ParseInt(req.ModelType, 10, 64)
	InOperation, _ := strconv.ParseInt(req.InOperation, 10, 64)
	pro := model.PropertyInfo{
		IndustryID:  req.IndustryID,
		Title:       req.Title,
		Nickname:    req.Nickname,
		Telephone:   req.Telephone,
		Image:       req.Image,
		CityCode:    req.CityCode,
		Idling:      req.Idling,
		InOperation: InOperation,
		TransferFee: TransferFee,
		Description: req.Description,
		ShopName:    req.ShopName,
		BusType:     BusType,
		ModelType:   ModelType,
		ExplicitTel: req.ExplicitTel,
		Tel1:        req.Tel1,
		Tel2:        req.Tel2,
		Protect:     req.Protect,
		QuotedPrice: QuotedPrice,
		Remake:      req.Remake,
		SourceInfo:  req.SourceInfo,
		MinArea:     req.MinArea,
		MaxArea:     req.MaxArea,
		MinRent:     req.MinRent,
		MaxRent:     req.MaxRent,
		Lots:        lots,
	}
	if err := pro.CreatePropertyInfo(); err != nil {
		res.Message = "发布求租求购失败"
		return
	}
	res.Status = true
	res.Message = "发布成功"
	return
}

// 修改求租求购
func EditQZQGProInfo(id int64, req entity.AddQZQGPropertyInfoRequest) (res entity.ResponseData) {
	proInfo := model.PropertyInfoScan{}
	proInfo.ID = id
	if err := proInfo.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
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
	if strings.Trim(req.Telephone, " ") == "" {
		res.Message = "联系手机不能为空"
		return
	}
	proInfos, count := model.QueryPropertyInfo(1, 0, map[string]interface{}{"telephone": req.Telephone})
	if count > 0 && proInfos[0].ID != id {
		res.Message = "此联系人已存在物业信息，请勿重复添加"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "联系手机格式不正确"
		return
	}
	city := model.City{
		Code: req.CityCode,
	}
	if err := city.QueryCitysByCode(); err != nil {
		res.Message = "城市代码不存在"
	}
	if !(req.BusType == "0" || req.BusType == "1" || req.BusType == "2") {
		req.BusType = "0"
	}
	if !(req.ModelType == "4" || req.ModelType == "5") {
		req.BusType = "4"
	}
	if req.MinArea <= 0 {
		res.Message = "最小面积必须大于0"
		return
	}
	if req.MaxArea <= 0 {
		res.Message = "最小面积必须大于0"
		return
	}
	if req.MinRent <= 0 {
		res.Message = "最小租金必须大于0"
		return
	}
	if req.MaxRent <= 0 {
		res.Message = "最小租金必须大于0"
		return
	}
	if req.MinArea > req.MaxArea {
		res.Message = "最小面积必须不大于最大面积"
		return
	}
	if req.MinRent > req.MaxRent {
		res.Message = "最小租金必须不大于最大租金"
		return
	}
	if req.Lots == "" {
		res.Message = "考虑地段至少选择一个"
		return
	}
	var lots []model.Lot
	arr := strings.Split(req.Lots, ",")
	for _, v := range arr {
		dis := model.District{
			Code: v,
		}
		if err := dis.QueryDistrictByCode(); err != nil {
			res.Message = "考虑地段区域不存在"
			return
		}
		lot := model.Lot{
			DistrictCode: dis.Code,
			DistrictName: dis.Name,
		}
		lots = append(lots, lot)
	}
	TransferFee, _ := strconv.ParseFloat(req.TransferFee, 64)
	QuotedPrice, _ := strconv.ParseFloat(req.QuotedPrice, 64)
	args := map[string]interface{}{
		"industry_id":  req.IndustryID,
		"title":        req.Title,
		"ickname":      req.Nickname,
		"relephone":    req.Telephone,
		"image":        req.Image,
		"city_code":    req.CityCode,
		"idling":       req.Idling,
		"in_operation": req.InOperation,
		"transfer_fee": TransferFee,
		"description":  req.Description,
		"shop_name":    req.ShopName,
		"bus_type":     req.BusType,
		"model_type":   req.ModelType,
		"explicit_tel": req.ExplicitTel,
		"tel1":         req.Tel1,
		"tel2":         req.Tel2,
		"protect":      req.Protect,
		"quoted_price": QuotedPrice,
		"remake":       req.Remake,
		"source_info":  req.SourceInfo,
		"min_area":     req.MinArea,
		"max_area":     req.MaxArea,
		"min_rent":     req.MinRent,
		"max_rent":     req.MaxRent,
		"lots":         lots,
	}
	if err := proInfo.EditPropertyInfoByID(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 修改物业信息
func EditProInfo(id int64, req entity.AddPropertyInfoRequest) (res entity.ResponseData) {
	info := model.PropertyInfoScan{}
	info.ID = id
	if err := info.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
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
	areaFlo, _ := strconv.ParseFloat(req.Area, 64)
	if areaFlo <= 0 {
		res.Message = "面积必须大于0"
		return
	}
	rentFlo, _ := strconv.ParseFloat(req.Rent, 64)
	if rentFlo <= 0 {
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
	pros, count := model.QueryPropertyInfo(30, 1, map[string]interface{}{"telephone": req.Telephone})
	if count > 0 && info.ID != pros[0].ID {
		res.Message = "此联系人已存在物业信息，请勿重复添加"
		return
	}
	if !utilsHelper.CheckTelFormat(req.Telephone) {
		res.Message = "联系手机格式不正确"
		return
	}
	if req.IndustryRanges == "" {
		res.Message = "适合经营至少选勾选一项"
		return
	}
	var industryRangeArr []model.IndustryRange
	arr := strings.Split(req.IndustryRanges, ",")
	for _, item := range arr {
		ind := model.Industry{}
		id, _ := strconv.ParseInt(item, 10, 64)
		ind.ID = id
		if err := ind.QueryIndustryByID(); err != nil {
			res.Message = fmt.Sprintf("经营业态ID:%s不存在", item)
			return
		}
		industryRange := model.IndustryRange{
			IndustryID:   ind.ID,
			IndustryName: ind.Name,
		}
		industryRangeArr = append(industryRangeArr, industryRange)
	}
	fmt.Println(industryRangeArr)
	storeType := model.StoreType{}
	storeType.ID = req.StoreTypeID
	if err := storeType.QueryStoreTypeByID(); err != nil {
		res.Message = "店铺类型不存在"
		return
	}
	rent := model.RentType{}
	if err := rent.QueryRentTypeByRent(rentFlo); err != nil {
		res.Message = "租金类型不存在"
		return
	}
	area := model.AreaType{}
	if err := area.QueryAreaTypeByArea(areaFlo); err != nil {
		res.Message = "面积类型不存在"
		return
	}
	transferFee, _ := strconv.ParseFloat(req.TransferFee, 64)
	quotedPrice, _ := strconv.ParseFloat(req.QuotedPrice, 64)
	args := map[string]interface{}{
		"industry_id":     req.IndustryID,
		"title":           req.Title,
		"nickname":        req.Nickname,
		"telephone":       req.Telephone,
		"shop_name":       req.ShopName,
		"image":           req.Image,
		"bus_type":        req.BusType,
		"model_type":      req.ModelType,
		"province_code":   req.ProvinceCode,
		"city_code":       req.CityCode,
		"district_code":   req.DistrictCode,
		"street_code":     req.StreetCode,
		"industry_ranges": industryRangeArr,
		"area_type_id":    area.ID,
		"rent_type_id":    rent.ID,
		"area":            areaFlo,
		"rent":            rentFlo,
		"address":         req.Address,
		"description":     req.Description,
		"idling":          req.Idling,
		"remake":          req.Remake,
		"transfer_fee":    transferFee,
		"explicit_tel":    req.ExplicitTel,
		"tel1":            req.Tel1,
		"tel2":            req.Tel2,
		"in_operation":    req.InOperation,
		"protect":         req.Protect,
		"store_type_id":   req.StoreTypeID,
		"quoted_price":    quotedPrice,
	}
	if err := info.EditPropertyInfoByID(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Message = "修改成功"
	res.Status = true
	return
}

// 删除物业信息
func DelProInfo(id int64) (res entity.ResponseData) {
	pro := model.PropertyInfo{}
	pro.ID = id
	if err := pro.DelProInfo(); err != nil {
		res.Message = "删除失败"
		return
	}
	res.Message = "删除成功"
	res.Status = true
	return
}

// 添加收款
func AddPayInfo(req entity.PayInfoRequest) (res entity.ResponseData) {
	info := model.PayInfo{
		ProInfoID: req.ProInfoID,
	}
	if err := info.QueryPayInfoByProInfoID(); err == nil {
		res.Message = "此物业已收款"
		return
	}
	fmt.Println(req)
	if req.Name == "" {
		res.Message = "收款人姓名不能为空"
		return
	}
	met := model.PayMethond{}
	met.ID = req.PayMethondID
	if err := met.QueryPayMethondByID(); err != nil {
		res.Message = "付款方式不存在"
		return
	}
	if req.PayTime == "" {
		res.Message = "收款时间不能为空"
		return
	}
	pro := model.PropertyInfoScan{}
	pro.ID = req.ProInfoID
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	user := model.User{}
	user.ID = req.PayeeID
	if err := user.QueryEmployeeById(); err != nil {
		res.Message = "业绩归属用户不存在"
		return
	}
	pay := model.PayInfo{
		Name:             req.Name,
		PayMethond:       met.Name,
		PayMethondID:     met.ID,
		PayStatus:        req.PayStatus,
		Payee:            user.Nickname,
		PayTime:          utilsHelper.StringToSTime(req.PayTime),
		PayeeID:          met.ID,
		ReceivableAmount: req.ReceivableAmount,
		ActualAmount:     req.ActualAmount,
		Invoice:          req.Invoice,
		Remake:           req.Remake,
		ProInfoID:        pro.ID,
		ProInfoNickname:  pro.Nickname,
		ProInfoTitle:     pro.Title,
	}
	if err := pay.AddPayInfo(); err != nil {
		res.Message = "收款失败"
		return
	}
	res.Status = true
	res.Message = "收款成功"
	return
}

// 修改收款
func EditPayInfoByProInfo(id int64, req entity.PayInfoRequestByProInfo) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	info := model.PayInfo{
		ProInfoID: id,
	}
	if err := info.QueryPayInfoByProInfoID(); err != nil {
		res.Message = "收款记录不存在，此物业暂未收款"
		return
	}
	if req.Name == "" {
		res.Message = "收款人姓名不能为空"
		return
	}
	met := model.PayMethond{}
	met.ID = req.PayMethondID
	if err := met.QueryPayMethondByID(); err != nil {
		res.Message = "付款方式不存在"
		return
	}
	if req.PayTime == "" {
		res.Message = "收款时间不能为空"
		return
	}
	user := model.User{}
	user.ID = req.PayeeID
	if err := user.QueryEmployeeById(); err != nil {
		res.Message = "业绩归属用户不存在"
		return
	}
	args := map[string]interface{}{
		"name":              req.Name,
		"pay_methond":       met.Name,
		"pay_methond_id":    met.ID,
		"pay_status":        req.PayStatus,
		"payee":             user.Nickname,
		"pay_time":          utilsHelper.StringToSTime(req.PayTime),
		"payee_id":          met.ID,
		"receivable_amount": req.ReceivableAmount,
		"actual_amount":     req.ActualAmount,
		"invoice":           req.Invoice,
		"remake":            req.Remake,
		"pro_info_id":       pro.ID,
		"pro_info_nickname": pro.Nickname,
		"pro_info_title":    pro.Title,
	}
	if err := info.EditPayInfo(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 修改收款
func EditPayInfo(id int64, req entity.PayInfoRequest) (res entity.ResponseData) {
	payInfo := model.PayInfo{}
	payInfo.ID = id
	if err := payInfo.QueryPayInfoByID(); err != nil {
		res.Message = "收款记录不存在"
		return
	}
	info := model.PayInfo{
		ProInfoID: req.ProInfoID,
	}
	if err := info.QueryPayInfoByProInfoID(); err != nil && id != info.ID {
		res.Message = "此物业已收款"
		return
	}
	if req.Name == "" {
		res.Message = "收款人姓名不能为空"
		return
	}
	met := model.PayMethond{}
	met.ID = req.PayMethondID
	if err := met.QueryPayMethondByID(); err != nil {
		res.Message = "付款方式不存在"
		return
	}
	if req.PayTime == "" {
		res.Message = "收款时间不能为空"
		return
	}
	pro := model.PropertyInfoScan{}
	pro.ID = req.ProInfoID
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	user := model.User{}
	user.ID = req.PayeeID
	if err := user.QueryEmployeeById(); err != nil {
		res.Message = "业绩归属用户不存在"
		return
	}
	args := map[string]interface{}{
		"name":              req.Name,
		"pay_methond":       met.Name,
		"pay_methond_id":    met.ID,
		"pay_status":        req.PayStatus,
		"payee":             user.Nickname,
		"pay_time":          utilsHelper.StringToSTime(req.PayTime),
		"payee_id":          met.ID,
		"receivable_amount": req.ReceivableAmount,
		"actual_amount":     req.ActualAmount,
		"invoice":           req.Invoice,
		"remake":            req.Remake,
		"pro_info_id":       pro.ID,
		"pro_info_nickname": pro.Nickname,
		"pro_info_title":    pro.Title,
	}
	if err := payInfo.EditPayInfo(args); err != nil {
		res.Message = "修改失败"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 收款详情（物业列表）
func QueryPayInfoByProInfo(id int64) (res entity.ResponseData) {
	pay := model.PayInfo{
		ProInfoID: id,
	}
	if err := pay.QueryPayInfoByProInfoID(); err != nil {
		res.Message = "未找到业绩信息"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{"pay_info": pay}
	return
}

// 收款详情（收款列表）
func QueryPayInfo(id int64) (res entity.ResponseData) {
	pay := model.PayInfo{}
	pay.ID = id
	if err := pay.QueryPayInfoByID(); err != nil {
		res.Message = "获取失败"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{"pay_info": pay}
	return
}

// 收款列表
func QueryPayInfos(pageSize int, page int, args map[string]interface{}) (res entity.ResponseData) {
	count, pay_infos := model.QueryPayInfo(pageSize, page, args)
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{"pay_infos": pay_infos, "count": count}
	return
}

// 添加付款方式
func AddPayMethond(req entity.PayMethondRequest) (res entity.ResponseData) {
	if req.Name == "" {
		res.Message = "付款方式不能为空"
		return
	}
	met := model.PayMethond{
		Name: req.Name,
		Card: req.Card,
	}
	if err := met.AddPayMethond(); err != nil {
		res.Message = "添加失败"
		return
	}
	res.Message = "添加成功"
	res.Status = true
	return
}

// 修改付款方式
func EditPayMethond(id int64, req entity.PayMethondRequest) (res entity.ResponseData) {
	met := model.PayMethond{}
	met.ID = id
	if err := met.QueryPayMethondByID(); err != nil {
		res.Message = "付款方式不存在"
		return
	}
	if req.Name == "" {
		res.Message = "付款方式不能为空"
		return
	}
	args := map[string]interface{}{
		"name": req.Name,
		"card": req.Card,
	}
	if err := met.EditPayMethond(args); err != nil {
		res.Message = "修改成功"
		return
	}
	res.Status = true
	res.Message = "修改成功"
	return
}

// 删除付款
func DelPayInfo(ids []int64) (res entity.ResponseData) {
	count := model.DelPayInfo(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 删除付款方式
func DelPayMethond(ids []int64) (res entity.ResponseData) {
	count := model.DelPayMethond(ids)
	if count == 0 {
		res.Message = "删除失败"
		return
	}
	res.Status = true
	res.Message = fmt.Sprintf("成功删除%v条，失败%v条", count, int64(len(ids))-count)
	return
}

// 付款方式详情
func QueryPayMethondByID(id int64) (res entity.ResponseData) {
	met := model.PayMethond{}
	met.ID = id
	if err := met.QueryPayMethondByID(); err != nil {
		res.Message = "付款方式不存在"
		return
	}
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{"pay_methond": met}
	return
}

// 付款方式列表
func QueryPayMethond() (res entity.ResponseData) {
	pay_methonds := model.QueryPayMethond()
	res.Status = true
	res.Message = "获取成功"
	res.Data = map[string]interface{}{"pay_methonds": pay_methonds, "count": len(pay_methonds)}
	return
}

// 查询用户菜单
func QueryUserMenu(token string) (res entity.ResponseData) {
	user := model.User{
		Token: token,
	}
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误"
		return
	}
	count, menus := model.QueryMenuByUser(user.UserInfo.RoleID)
	res.Status = true
	res.Data = map[string]interface{}{"menus": menus, "count": count}
	return
}

// 物业成功
func EditProInfoSuccess(id int64) (res entity.ResponseData) {
	pro := model.PropertyInfoScan{}
	pro.ID = id
	if err := pro.QueryPropertyInfoByID(); err != nil {
		res.Message = "物业信息不存在"
		return
	}
	args := map[string]interface{}{"status": true}
	if err := pro.EditPropertyInfoByID(args); err != nil {
		res.Message = "操作失败"
		return
	}
	res.Message = "操作成功"
	res.Status = true
	return
}
