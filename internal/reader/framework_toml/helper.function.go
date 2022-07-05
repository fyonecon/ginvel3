package framework_toml

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// Log 记录一般日志，定期自动删除
func Log(_txt string, _ip string) {
	// 读取项目绝对路径根目录
	mainDir, _ := os.Getwd()
	mainDirectory := mainDir + "/"
	storagePath := mainDirectory + "storage/"

	// 创建文件夹
	filepath := storagePath+"log_frame/ok/"
	dateFile := GetTimeDate("Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	if len(_ip) == 0 {
		_ip = ServerIPv4()[0]
	}

	// 文件信息
	fileName := "log_" + GetTimeDate("Y-m-d") + ".log"
	filePath := saveFilepath
	file, err := os.OpenFile(filePath+fileName, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
	if err != nil {
		log.Println("文件创建失败", filePath, err)
		//panic(err)
		return
	}
	// 延迟关闭文件
	defer file.Close()
	//
	date := GetTimeDate("Y-m-d H:i:s")
	txt := "【" + date + "】" + _ip + "：" + _txt
	// 写入文件内容
	_, err = file.WriteString(txt + "\n")
	if err != nil {
		log.Println("文件写入失败", filePath, err)
		return
	}

	// 删除老文件夹
	_oldMonth := -24
	delFile := time.Now().AddDate(0, _oldMonth, -1).Format("20060102")
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("老文件夹删除失败=", err1)
	}

}

// ServerIPv4 获取本机IP
func ServerIPv4() []string {
	var ipArray []string // IP数组

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						//fmt.Println(ipnet.IP.String())
						ipArray = append(ipArray, ipnet.IP.String())
					}
				}
			}
		}
	}

	return ipArray
}

// HasFile 判断文件或文件夹是否存在
func HasFile(filePath string) (bool, string) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, filePath
	}else {
		return false, "FileChecker:::NotFound " + filePath
	}
}

// IntToString int转string
func IntToString(_int int64) string {
	_str := strconv.FormatInt(_int,10)
	return _str
}

// GetTimeDate 获取日期时间戳，s
// Y年m月d号 H:i:s.MS.NS 星期W
func GetTimeDate(_format string) (date string) {
	if len(_format) == 0 {
		_format = "YmdHisMS"
	}
	date = _format

	// 时区
	//timeZone, _ := time.LoadLocation(FrameInfo["timezone"])
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timer := time.Now().In(timeZone)

	var year int64 = int64(timer.Year())
	var month int64 = int64(timer.Month())
	var day int64 = int64(timer.Day())
	var hour int64 = int64(timer.Hour())
	var minute int64 = int64(timer.Minute())
	var second int64 = int64(timer.Second())
	var week int64 = int64(timer.Weekday())
	var ms int64 = int64(timer.UnixNano() / 1e6)
	var ns int64 = int64(timer.UnixNano() / 1e9)
	msTmp := IntToString(int64(math.Floor(float64(ms/1000))))
	nsTmp := IntToString(int64(math.Floor(float64(ns/1000000))))

	var _year string
	var _month string
	var _day string
	var _hour string
	var _minute string
	var _second string
	var _week string // 英文星期
	var _Week string // 中文星期
	var _ms string // 毫秒
	var _ns string // 纳秒

	_year = IntToString(year)
	if month < 10 {
		_month = IntToString(month)
		_month = "0" + _month
	}else {
		_month = IntToString(month)
	}
	if day < 10 {
		_day = IntToString(day)
		_day = "0" + _day
	}else {
		_day = IntToString(day)
	}
	if hour < 10 {
		_hour = IntToString(hour)
		_hour = "0" + _hour
	}else {
		_hour = IntToString(hour)
	}
	if minute < 10 {
		_minute = IntToString(minute)
		_minute = "0" + _minute
	}else {
		_minute = IntToString(minute)
	}
	if second < 10 {
		_second = IntToString(second)
		_second = "0" + _second
	}else {
		_second = IntToString(second)
	}
	_week = IntToString(week)
	WeekZh := [...]string{"日", "一", "二", "三", "四", "五", "六"} // 默认从"日"开始
	_Week = WeekZh[week]
	_ms = strings.Replace(IntToString(ms), msTmp, "", -1)
	_ns = strings.Replace(IntToString(ns), nsTmp, "", -1)

	// 替换关键词
	date = strings.Replace(date, "MS", _ms, -1)
	date = strings.Replace(date, "NS", _ns, -1)
	date = strings.Replace(date, "Y", _year, -1)
	date = strings.Replace(date, "m", _month, -1)
	date = strings.Replace(date, "d", _day, -1)
	date = strings.Replace(date, "H", _hour, -1)
	date = strings.Replace(date, "i", _minute, -1)
	date = strings.Replace(date, "s", _second, -1)
	date = strings.Replace(date, "W", _Week, -1)
	date = strings.Replace(date, "w", _week, -1)

	return
}

// DownloadFile 下载文件，返回新文件绝对地址
func DownloadFile(fileUrl string, savePath string, newFilename string) ([]interface{}, []interface{}) {

	var file string = savePath + newFilename
	var fileInfo []interface{} // 文件信息：[文件全名包括后缀，新绝对地址，原链接，文件类型，文件大小KB]
	var errorArray []interface{} // 报错信息
	var size float64 // 文件大小
	var ext string // 文件后缀

	// 默认值
	if len(fileUrl) == 0 {
		errorArray = append(errorArray, "文件链接为空")
	}
	if len(savePath) == 0 {
		errorArray = append(errorArray, "文件保存的绝对地址文件夹为空")
	}
	if len(newFilename) == 0 {
		errorArray = append(errorArray, "文件全名为空")
	}

	// 读
	resp, err := http.Get(fileUrl)
	if err != nil {
		errorArray = append(errorArray, err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(file)
	if err != nil {
		errorArray = append(errorArray, err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		errorArray = append(errorArray, err)
	}

	ext = path.Ext(file)
	if len(ext) == 0 {
		errorArray = append(errorArray, "newFilename无后缀")
	}
	
	fmt.Println(path.Dir(file))

	fileStat, err:=os.Stat(file)
	if err != nil {
		errorArray = append(errorArray, err)
	}else {
		_size := fileStat.Size()
		size = float64(_size/1024)
	}

	// 正确的文件信息
	fileInfo = []interface{}{newFilename, file, fileUrl, ext, size}

	log.Println("下载文件的信息：", fileInfo, errorArray)

	return fileInfo, errorArray
}

// RandRange 获取指定范围内的可变随机整数数，正负都行。[a, b]
func RandRange(_min int64, _max int64) int64 {
	var _rand int64
	if _min >= _max {
		_rand = 0
	}else {
		rand.Seed(time.Now().UnixNano())
		_rand = rand.Int63n(_max - _min) + _min
	}
	return _rand
}

// ValueInterfaceToString interface转string，非map[string]interface{}
func ValueInterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}