package kits

import (
	"fmt"
	"ginvel/common/helper"
	"ginvel/internal/providers"
	"github.com/gin-gonic/gin"
	"time"
)

// UpdateIP 记录IP，
// 返回：总PV、24h内pv
func UpdateIP(ctx *gin.Context, ip string) (int64, int64) {
	//
	allPVKey := "ip-pv-any-"+ip
	allPV, err := providers.RedisDb.Get(ctx, allPVKey).Result()
	if err != nil {
		allPV = "0"
	}
	_allPV := helper.StringToInt(allPV)
	newAllPV := helper.IntToString(_allPV + 1)
	fmt.Println("UpdateIP", allPVKey, allPV, _allPV, newAllPV)
	//
	err3 := providers.RedisDb.Set(ctx, allPVKey, newAllPV, 0).Err() // 永不过期
	if err3 != nil {
		fmt.Println(err3)
	}

	//
	hour24Key := "ip-pv-h24-"+ip
	hour24PV, err := providers.RedisDb.Get(ctx, hour24Key).Result()
	if err != nil {
		hour24PV = "0"
	}
	_hour24PV := helper.StringToInt(hour24PV)
	newHour24PV := helper.IntToString(_hour24PV + 1)
	fmt.Println("UpdateIP", hour24Key, hour24PV, _hour24PV, newHour24PV)
	//
	err4 := providers.RedisDb.Set(ctx, hour24Key, newHour24PV, time.Hour * 24).Err() // 24小时后过期
	if err4 != nil {
		fmt.Println(err4)
	}

	return _allPV, _hour24PV
}


// SearchIP 查询IP的PV
// 返回：总PV、24h内pv
func SearchIP(ctx *gin.Context, ip string) (int64, int64){
	//
	allPVKey := "pv-all-"+ip
	allPV, err := providers.RedisDb.Get(ctx, allPVKey).Result()
	if err != nil {
		allPV = "0"
	}
	_allPV := helper.StringToInt(allPV)
	newAllPV := helper.IntToString(_allPV + 1)
	fmt.Println("SearchIP=", allPVKey, allPV, _allPV, newAllPV)

	//
	hour24Key := "pv-hour24-"+ip
	hour24PV, err := providers.RedisDb.Get(ctx, hour24Key).Result()
	if err != nil {
		hour24PV = "0"
	}
	_hour24PV := helper.StringToInt(hour24PV)
	newHour24PV := helper.IntToString(_hour24PV + 1)
	fmt.Println("SearchIP=", hour24Key, hour24PV, _hour24PV, newHour24PV)

	return _allPV, _hour24PV
}