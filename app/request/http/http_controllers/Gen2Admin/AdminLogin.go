package Gen2Admin

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	"ginvel/internal/providers"
	"github.com/gin-gonic/gin"
)

// AdminLogin 用户登录
func AdminLogin(ctx *gin.Context)  {
	var state int64
	var msg string
	var err error
	var tokenNum int = 4 // 可以同时登录几台设备或客户端
	var loginToken string

	// 参数
	appClass := kits.Input(ctx, "app_class")
	loginName := kits.Input(ctx, "login_name")
	loginPwd := kits.Input(ctx, "login_pwd")
	loginCaptcha := kits.Input(ctx, "login_captcha")
	ua := kits.Input(ctx, "ua")
	//url := kits.Input(ctx, "url")

	if len(loginPwd) >= 6 && len(loginPwd) <= 200 {
		loginPwd = helper.AdminPwd(loginPwd)
	}else {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "输入参数不符合规范",
			"content": map[string]interface{}{
				"login_name": loginName,
				"login_pwd": loginPwd,
				"login_captcha": loginCaptcha,
			},
		})
		return
	}

	// 用户是否存在，然后对比两个密码是否匹配
	type AdminHasKeys struct {
		AdminID int64 `json:"admin_id"`
		LoginName string `json:"login_name"`
		LoginPWD string `json:"login_pwd"`
		LoginLevel string `json:"login_level"`
	}
	var hasData AdminHasKeys  // 或hasData := AdminHasKeys{}
	_whereMap := map[string]interface{}{}
	_whereMap["login_name"] = loginName
	_notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [2, 3, 4]
	}
	hasRes := providers.GDB.Table("gv_admin").Not(_notMap).Where(_whereMap).First(&hasData)
	err = hasRes.Error
	if err != nil {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "登录失败，帐户信息不匹配，-3",
			"content": map[string]interface{}{
				"login_name": loginName,
				"login_pwd": loginPwd,
				"err": err,
			},
		})
		return
	}
	adminID := hasData.AdminID
	adminPwd := hasData.LoginPWD
	loginLevel := hasData.LoginLevel

	if loginPwd == adminPwd { // 用户+密码 正确

		// 结果集
		type ListAdminTokenKeys struct {
			//ID int64 // 要返回的主键值
			AdminTokenID int64 `json:"admin_token_id"`
			AdminID int64 `json:"admin_id"`
			LoginName string `json:"login_name"`
			LoginLevel string `json:"login_level"`
			UpdateTime  string `json:"update_time"`
			UpdateIP  string `json:"update_ip"`
		}
		var res []ListAdminTokenKeys

		whereMap := map[string]interface{}{}
		whereMap["admin_id"] = adminID

		// 数据列表
		list := providers.GDB.Table("gv_admin_token").Where(whereMap).Order("update_time asc").Limit(tokenNum).Scan(&res)
		err = list.Error

		updateTime := helper.GetTimeDate("YmdHis")
		updateIP := ctx.ClientIP()
		// 新token
		newToken := CreateToken(helper.IntToString(adminID), loginName)

		if len(res) >= tokenNum { // 替换
			oldestAdminTokenID := res[0].AdminTokenID
			data := map[string]interface{}{
				"login_token": newToken,
				"update_time": updateTime,
				"update_ip": updateIP,
				"ua": ua,
				"app_class": appClass,
			}
			__whereMap := map[string]interface{}{}
			__whereMap["admin_token_id"] = oldestAdminTokenID
			// 操作数据库
			res := providers.GDB.Table("gv_admin_token").Where(__whereMap).Updates(data)
			err = res.Error
			row := res.RowsAffected
			if row > 0 {
				state = 1
				msg = "登录成功，2"
				loginToken = newToken
			}else {
				state = 0
				msg = "登录失败，系统内部信息不匹配，-2"
			}
		}else{ // 新增
			// MySQL数据格式
			type AdminTokenKeys struct {
				ID int64 // 要返回的主键值
				AdminID  int64 `json:"admin_id"`
				AppClass  string `json:"app_class"`
				LoginToken string `json:"login_token"`
				UpdateTime  string `json:"update_time"`
				UpdateIP  string `json:"update_ip"`
			}
			// 新数据
			data := AdminTokenKeys{
				AdminID: adminID,
				AppClass: appClass,
				LoginToken: newToken,
				UpdateTime: updateTime,
				UpdateIP:   updateIP,
			}
			res := providers.GDB.Table("gv_admin_token").Create(&data)
			adminTokenID := data.ID
			err = res.Error
			if adminTokenID > 0{
				state = 1
				msg = "登录成功，1"
				loginToken = newToken
			}else {
				state = 0
				msg = "登录失败，数据不能写入，-3"
			}
		}
	}else {
		state = 0
		msg = "登录失败，帐户信息不匹配，-1"
	}

	// 接口返回
	back := map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": map[string]interface{}{
			"err": err,
			"app_class": appClass,
			"login_id": adminID,
			"login_name": loginName,
			"login_token": helper.EncodeURL(loginToken),
			"login_pwd": loginPwd,
			"level_info": map[string]interface{}{
				"login_level": loginLevel,
				"login_position": ShowAdminLevelPosition(helper.StringToInt(loginLevel)),
			},
		},
	}
	ctx.JSONP(200, back)
	return
}

// CheckLoginToken 保持登录状态的检测（并记录日志）
// 统一在拦截器里面做校验
func CheckLoginToken(ctx *gin.Context)  {
	appClass := kits.Input(ctx, "app_class")
	loginID := kits.Input(ctx, "login_id")
	loginName := kits.Input(ctx, "login_name")
	loginLevel := "已登录"

	// 接口返回
	back := map[string]interface{}{
		"state": 1,
		"msg": "Token可用",
		"content": map[string]interface{}{
			"app_class": appClass,
			"login_id": loginID,
			"login_name": loginName,
			"login_level": loginLevel,
		},
	}
	ctx.JSONP(200, back)
	return
}
