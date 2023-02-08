package Gen2Admin

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	"ginvel/internal/providers"
	"github.com/gin-gonic/gin"
)

// ListAdmin 用户列表
func ListAdmin(ctx *gin.Context)  {
	var state int64
	var msg string
	var err error

	_adminID := kits.Input(ctx, "admin_id")
	adminID := helper.StringToInt(_adminID)
	loginName := kits.Input(ctx, "login_name")
	_page := kits.Input(ctx, "page")
	page := helper.StringToInt(_page)

	// 处理分页
	var limit int64 = helper.Paging["limit"]
	var offset int64 = 0 // 本页从第几个开始
	if page <= 0 { page = 1 }
	page = page - 1
	offset = limit*page

	// 结果集
	type ListAdminKeys struct {
		//ID int64 // 要返回的主键值
		AdminID int64 `json:"admin_id"`
		LoginName string `json:"login_name"`
		CreateTime  string `json:"create_time"`
		UpdateIP  string `json:"update_ip"`
	}
	var res []ListAdminKeys
	var total int64

	// 多条件查询
	whereMap := map[string]interface{}{}
	notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [0, 2]
	}

	// 数据列表
	list := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap)
	// 数据总数
	count := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap)

	// 加入非"="多条件查询
	if len(loginName) != 0 {
		list = list.Where("login_name LIKE ?", "%" + loginName + "%")
		count = count.Where("login_name LIKE ?", "%" + loginName + "%")
	}
	if adminID != 0 {
		list = list.Where("admin_id LIKE ?", "%" + helper.IntToString(adminID))
		count = count.Where("admin_id LIKE ?", "%" + helper.IntToString(adminID))
	}

	// 完成其他剩余条件
	list.Order("admin_id desc").Limit(int(limit)).Offset(int(offset)).Scan(&res)
	count.Order("admin_id desc").Count(&total)

	err = list.Error
	if len(res) > 0 {
		state = 1
		msg = "查询完成"
	}else {
		state = 0
		msg = "查询无数据"
	}

	// 遍历struct数据，修整数据
	for i := 0; i < len(res); i++ {
		theCreateTime := res[i].CreateTime
		newCreateTime := helper.DateToDate(theCreateTime)
		res[i].CreateTime = newCreateTime
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"whereMap": whereMap,
		"err": err,
	}

	// 分页数据
	paging := helper.MakePaging(total, limit, page)
	// 返回数据
	back := map[string]interface{}{
		"state": state,
		"msg": msg,
		"paging": paging,
		"test_data": testData,
		"content": res,
	}

	ctx.JSONP(200, back)
	return
}

// ThatAdmin 此用户的详细信息
func ThatAdmin(ctx *gin.Context)  {
	var state int64
	var msg string
	var err error

	loginName := kits.Input(ctx, "login_name")

	// 结果集
	type ThatAdminKeys struct {
		//ID int64 // 要返回的主键值
		AdminID int64 `json:"admin_id"`
		LoginName string `json:"login_name"`
		CreateTime  string `json:"create_time"`
		UpdateIP  string `json:"update_ip"`
	}
	var hasData ThatAdminKeys // 或hasData := ThatAdminKeys{}

	// 多查询条件
	whereMap := map[string]interface{}{}
	whereMap["login_name"] = loginName
	notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [0, 2]
	}
	res := providers.GDB.Table("gv_admin").Where(whereMap).Not(notMap).First(&hasData)
	err = res.Error
	if hasData.AdminID != 0 {
		state = 1
		msg = "查询完成"
	}else {
		state = 0
		msg = "查询无数据"
	}

	// 访问结构体并改变成员变量的值
	createTime := hasData.CreateTime
	createTime = helper.DateToDate(createTime)
	hasData.CreateTime = createTime

	// 返回一些测试数据
	testData := map[string]interface{}{
		"whereMap": whereMap,
		"err": err,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   hasData,
	})
	return
}

// AddAdmin 新增用户
func AddAdmin(ctx *gin.Context) {
	var state int64
	var msg string
	var id int64
	var err error

	// 参数
	loginName := kits.Input(ctx, "login_name")
	_loginPwd := kits.Input(ctx, "login_pwd")
	loginPwd := ""
	if len(_loginPwd) >= 6 && len(_loginPwd) <= 200 {
		loginPwd = helper.AdminPwd(_loginPwd)
	}else {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "输入参数不符合规范",
		})
		return
	}
	createTime := helper.GetTimeDate("YmdHis")
	updateIP :=ctx.ClientIP()

	// 结果集
	type AdminHasKeys struct {
		AdminID int64 `json:"admin_id"`
		LoginName string `json:"login_name"`
		CreateTime  string `json:"create_time"`
	}
	var hasData AdminHasKeys  // 或hasData := AdminHasKeys{}

	// 查询条件
	whereMap := map[string]interface{}{}
	whereMap["login_name"] = loginName
	notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [2, 3, 4]
	}
	hasRes := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap).First(&hasData)
	err = hasRes.Error
	if hasData.LoginName != ""{
		state = 0
		msg = "登录名已存在"
	}else {
		// 按账号优先级
		if len(loginName) >= 8 {

			// MySQL数据格式
			type AdminKeys struct {
				ID int64 // 要返回的主键值
				LoginName string `json:"login_name"`
				LoginPwd  string `json:"login_pwd"`
				CreateTime  string `json:"create_time"`
				UpdateIP  string `json:"update_ip"`
			}
			// 新数据
			data := AdminKeys{
				LoginName: loginName,
				LoginPwd: loginPwd,
				CreateTime: createTime,
				UpdateIP: updateIP,
			}

			// 操作数据库
			res := providers.GDB.Table("gv_admin").Select( "login_name", "login_pwd", "create_time", "update_ip").Create(&data)
			id = data.ID
			err = res.Error
			if err == nil{
				state = 1
				msg = "创建成功"
			}else {
				state = 0
				msg = "创建失败"
			}

		}else {
			state = 0
			msg = "账号信息不符合规范"
		}

	}

	// 接口返回
	back := map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": map[string]interface{}{
			"id": id,
			"err": err,
			"map": whereMap,
			"has": hasData,
			"_login_pwd": _loginPwd,
			"login_pwd": loginPwd,
			"login_name": loginName,
			"time": createTime,
			"ip": updateIP,
		},
	}
	ctx.JSONP(200, back)
	return
}

// UpdateAdmin 更新用户信息
func UpdateAdmin(ctx *gin.Context) {
	var state int64
	var msg string
	var err error
	var row int64

	// 参数
	adminId := kits.Input(ctx, "admin_id")
	loginName := kits.Input(ctx, "login_name")
	_loginPwd := kits.Input(ctx, "login_pwd")
	loginPwd := ""
	if len(_loginPwd) >= 6 && len(_loginPwd) <= 200 {
		loginPwd = helper.AdminPwd(_loginPwd)
	}else {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "输入参数不符合规范",
		})
		return
	}
	updateTime := helper.GetTimeDate("YmdHis")
	updateIP := ctx.ClientIP()

	// 确保login_name数据库唯一
	type AdminHasKeys struct {
		AdminID int64 `json:"admin_id"`
		LoginName string `json:"login_name"`
		CreateTime  string `json:"create_time"`
	}
	var hasData AdminHasKeys  // 或hasData := AdminHasKeys{}

	// 查询条件
	whereMap := map[string]interface{}{}
	whereMap["login_name"] = loginName
	notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [2, 3, 4]
	}
	hasRes := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap).First(&hasData)
	err = hasRes.Error
	if hasData.AdminID != 0 && hasData.AdminID != helper.StringToInt(adminId){
		state = 0
		msg = "登录名已存在"
	}else {
		if len(loginName) >= 8 {
			// 新数据
			data := map[string]interface{}{}
			if  len(_loginPwd) >= 8 && len(_loginPwd) <= 20 { // 更新pwd字段
				data = map[string]interface{}{
					"login_name": loginName,
					"login_pwd": loginPwd,
					"update_time": updateTime,
					"update_ip": updateIP,
				}
			}else {
				data = map[string]interface{}{
					"login_name": loginName,
					"update_time": updateTime,
					"update_ip": updateIP,
				}
			}

			// 查询条件
			whereMap := map[string]interface{}{}
			whereMap["admin_id"] = adminId
			notMap := map[string]interface{}{
				"state": []int64{0, 2}, // state <> [2, 3, 4]
			}

			// 操作数据库
			res := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap).Updates(data)
			err = res.Error
			row = res.RowsAffected
			if row > 0 {
				state = 1
				msg = "更新成功"
			}else {
				state = 0
				msg = "更新失败，可能是数据不存在"
			}
		}else {
			state = 0
			msg = "账号信息不符合规范"
		}
	}

	// 接口返回
	back := map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": map[string]interface{}{
			"row": row,
			"err": err,
			"🆔": adminId,
			"🆔2": hasData.AdminID,
			"_login_pwd": _loginPwd,
			"login_pwd": loginPwd,
			"login_name": loginName,
			"date": updateTime,
			"ip": updateIP,
		},
	}
	ctx.JSONP(200, back)
	return
}

// DelAdmin 删除用户信息，原则：不做实际数据删除
func DelAdmin(ctx *gin.Context) {
	var state int64
	var msg string
	var err error
	var row int64

	// 参数
	adminId := kits.Input(ctx, "admin_id")
	loginName := kits.Input(ctx, "login_name")

	updateTime := helper.GetTimeDate("YmdHis")
	updateIP :=ctx.ClientIP()

	// 新数据
	data := map[string]interface{}{
		"state": 2,
		"update_time": updateTime,
		"update_ip": updateIP,
	}

	// 查询条件
	whereMap := map[string]interface{}{}
	whereMap["admin_id"] = adminId
	whereMap["login_name"] = loginName
	notMap := map[string]interface{}{
		"state": []int64{0, 2}, // state <> [2, 3, 4]
	}

	// 操作数据库
	res := providers.GDB.Table("gv_admin").Not(notMap).Where(whereMap).Updates(data)
	err = res.Error
	row = res.RowsAffected
	if row > 0 {
		state = 1
		msg = "已删除"
	}else {
		state = 0
		msg = "删除失败，可能是数据不存在"
	}

	// 接口返回
	back := map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": map[string]interface{}{
			"err": err,
			"date": updateTime,
			"ip": updateIP,
		},
	}
	ctx.JSONP(200, back)
	return
}

// LevelAdmin 等级列表
func LevelAdmin(ctx *gin.Context)  {
	// 接口返回
	back := map[string]interface{}{
		"state": 1,
		"msg": "等级列表获取完成",
		"content": AdminLevelInfo(),
	}
	ctx.JSONP(200, back)
	return
}
