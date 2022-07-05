package Gen2User

// CheckAppClass 校验app_class值
func CheckAppClass(appClass string) bool {
	return true
}

func CheckToken(userId string, userName string, userToken string, appClass string) map[string]interface{}{
	var state int64
	var msg string
	//var footTime int64 = 30*24*60*60*1000 // ms，30 day

	return map[string]interface{}{
		"state": state,
		"msg": msg,
		"content": []interface{}{},
	}
}