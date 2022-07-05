package Gen2Admin

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	"ginvel/internal/providers"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 生成和验证验证码

// MakeCaptcha 生成验证码条件
func MakeCaptcha(ctx *gin.Context)  {
	var state int64
	var msg string
	var id int64
	var err error

	var uuid string = kits.Input(ctx, "UUID") // 前端生成
	var captcha int64 = helper.RandRange(1000, 9999)
	var createTime string = helper.IntToString(helper.GetTimeS())

	// 验证是否重复
	// 结果集
	type HasKeys struct {
		CaptchaID int64 `json:"captcha_id"`
		UUID string `json:"uuid"`
		Captcha int64 `json:"captcha"`
		UseNum int64 `json:"use_num"`
		CreateTime  string `json:"create_time"`
	}
	var hasData HasKeys  // 或hasData := AdminHasKeys{}
	// 查询条件
	whereMap := map[string]interface{}{}
	whereMap["uuid"] = uuid
	hasRes := providers.GDB.Table("gv_admin").Where(whereMap).First(&hasData)
	err = hasRes.Error
	if hasData.UUID != ""{
		state = 0
		msg = "uuid已存在"
		id = hasData.CaptchaID
	}else {
		// 保存到数据库
		// MySQL数据格式
		type AdminKeys struct {
			ID int64 // 要返回的主键值
			UUID string `json:"uuid"`
			Captcha  string `json:"captcha"`
			CreateTime  string `json:"create_time"`
		}
		// 新数据
		data := AdminKeys{
			UUID:       uuid,
			Captcha: helper.IntToString(captcha),
			CreateTime: createTime,
		}

		// 操作数据库
		res := providers.GDB.Table("gv_captcha").Create(&data)
		id = data.ID
		err = res.Error
		if err == nil{
			state = 1
			msg = "已生成"
		}else {
			state = 0
			msg = "未生成"
		}
	}

	if state == 1 {
		kits.MakeCaptcha(ctx, strconv.FormatInt(captcha, 10))
	}else {
		kits.MakeCaptcha(ctx, "null")
		kits.Log(msg, "captcha"+helper.IntToString(id))
	}

	return
}

// CheckCaptcha 验证验证码
func CheckCaptcha(uuid string, captcha string) (state int64, msg string, id int64, err error) {
	//var state int64
	//var msg string
	//var id int64
	//var err error
	var maxUseNum int64 = 2 // 每个验证码最大使用次数
	var foot int64 = 60 // s，验证码有效期

	//var uuid string = kits.Input(ctx, "UUID") // 前端生成，并复用
	//var captcha string = kits.Input(ctx, "captcha")

	// 验证是否在用
	// 结果集
	type HasKeys struct {
		CaptchaID int64 `json:"captcha_id"`
		UUID string `json:"uuid"`
		Captcha int64 `json:"captcha"`
		UseNum int64 `json:"use_num"`
		CreateTime  string `json:"create_time"`
	}
	var hasData HasKeys  // 或hasData := AdminHasKeys{}
	// 查询条件
	whereMap := map[string]interface{}{}
	whereMap["uuid"] = uuid
	hasRes := providers.GDB.Table("gv_admin").Where(whereMap).First(&hasData)
	err = hasRes.Error
	if hasData.UUID != ""{ // 存在
		id = hasData.CaptchaID
		useNum := hasData.UseNum
		createTime := helper.StringToInt(hasData.CreateTime)

		if maxUseNum <= useNum { // 失效
			state = 0
			msg = "验证码已失效，-2"
		}else {
			// 更新查询次数
			data := map[string]interface{}{
				"use_num": useNum + 1,
			}
			// 查询条件
			whereMap := map[string]interface{}{}
			whereMap["uuid"] = uuid
			// 操作数据库
			res := providers.GDB.Table("gv_admin").Where(whereMap).Updates(data)
			err = res.Error
			row := res.RowsAffected
			if row > 0 {

				// 判断时间
				nowTime := helper.GetTimeS()
				_foot := nowTime - createTime
				if _foot > foot {
					state = 0
					msg = "验证码已失效，-1"
				}else{
					state = 1
					msg = "已更新查询次数"
				}

			}else {
				state = 0
				msg = "查询次数失败，可能是数据不存在"
			}
		}

	}

	return state, msg, id, err
}