package commonService

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/utilsHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
)

// 生成验证码
func VerificationCode(tel string) (res entity.ResponseData) {
	if tel == "" {
		res.Message = "手机号码不能为空"
		return
	}
	if !utilsHelper.CheckTelFormat(tel) {
		res.Message = "手机号码格式不正确"
		return
	}
	code := utilsHelper.GenValidateCode(6)
	if code == "" {
		res.Message = "获取验证码失败"
		return
	}
	v := model.Verificationcode{
		Tel:        tel,
		CreateTime: utilsHelper.GetTimestamp(),
		Code:       code,
	}
	if err := v.AddVerificationcode(); err != nil {
		res.Message = "获取验证码失败"
		return
	}
	// 短信通知接口
	data := make(map[string]interface{})
	data["code"] = code
	res.Data = data
	res.Message = "获取验证码成功"
	res.Status = true
	return
}

// 获取用户信息 by token
func QueryUserByToken(token string) (user model.User, res entity.ResponseData) {
	if token == "" {
		res.Message = "token不能为空"
		return
	}
	user.Token = token
	if err := user.QueryByToken(); err != nil {
		res.Message = "token错误，用户信息不存在"
		return
	}
	res.Status = true
	res.Message = "查询成功"
	return
}
