package Gen2Admin

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	"strings"
)

// CheckAppClass 校验app_class值
func CheckAppClass(appClass string) bool {
	return true
}

// ShowAdminLevelPosition 根据等级返回身份描述
func ShowAdminLevelPosition(level int64) (position string) {
	for _, info := range AdminLevelInfo(){
		_theLevel := info["level"]
		_thePosition := info["position"]
		var theLevel int64 = helper.InterfaceToInt(_theLevel)
		var thePosition string = helper.InterfaceToString(_thePosition)
		if level == theLevel {
			position = thePosition
			break
		}
	}
	return position
}

// AdminLevelInfo Admin等级及其解释
func AdminLevelInfo() []map[string]interface{} {
	return []map[string]interface{}{
		map[string]interface{}{
			"level": 10,
			"position": "超管",
		},
		map[string]interface{}{
			"level": 12,
			"position": "组长",
		},
		map[string]interface{}{
			"level": 14,
			"position": "成员",
		},
		map[string]interface{}{
			"level": 16,
			"position": "游客",
		},
		map[string]interface{}{
			"level": 20,
			"position": "（违规被封）",
		},
		map[string]interface{}{
			"level": 22,
			"position": "（帐户过期）",
		},
		map[string]interface{}{
			"level": 24,
			"position": "（禁言状态）",
		},
	}
}

// CreateToken 创建token
func CreateToken(adminId string, loginName string) string {
	nowTime := helper.GetTimeDate("YmdHisMS")
	return kits.Encode(nowTime+"#@"+adminId+"#@"+loginName+"#@"+helper.RandString(helper.RandRange(6, 8)), "gKz0Z2lI")
}

// CheckToken 解析和检测token
func CheckToken(adminId string, loginName string, loginToken string, appClass string) map[string]interface{} {
	var state int64
	var msg string
	var footTime int64 = 30*24*60*60*1000 // ms，30 day

	_nowTime := helper.GetTimeDate("YmdHisMS")
	theToken := kits.Decode(loginToken, "gKz0Z2lI")
	arrayToken := strings.Split(theToken, "#@")
	if len(arrayToken) >= 4 {
		_theOldTime := arrayToken[0]
		theAdminId := arrayToken[1]
		theLoginName := arrayToken[2]

		if adminId == theAdminId && loginName == theLoginName {
			nowTime := helper.DateToTimeS(_nowTime, "YmdHisMS")
			theOldTime := helper.DateToTimeS(_theOldTime, "YmdHisMS")
			foot := nowTime-theOldTime
			if foot>footTime && foot>0 { // 过期
				state = 0
				msg = "token不可用，-1"
			}else { // 可用
				state = 1
				msg = "token可用，1"
			}
		}else {
			state = 0
			msg = "token不可用，-3"
		}
	}else {
		state = 0
		msg = "token不可用，-2"
	}

	return map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": []interface{}{
			arrayToken,
			adminId,
			loginName,
			loginToken,
			appClass,
		},
	}
}
