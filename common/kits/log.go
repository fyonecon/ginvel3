package kits

import (
	"fmt"
	"ginvel/common/helper"

	"log"
	"os"
	"time"
)

// Log 记录一般日志
func Log(_txt string, _ip string) {
	filename := "log_" + helper.GetTimeDate("m-d_H")
	WriteLog("sys_log/log/", _txt, _ip, filename)
}

// Error 记录错误日志
func Error(_txt string, _ip string) {
	filename := "error_" + helper.GetTimeDate("m-d_H")
	WriteLog("sys_log/error/", _txt, _ip, filename)
}

// Limit 记录限流日志
func Limit(_txt string, _ip string) {
	filename := "limit_" + helper.GetTimeDate("m-d_H")
	WriteLog("sys_log/limit/", _txt, _ip, filename)
}

// WriteLog 写入日志信息接口
// _filepath默认的父级目录是storage，取值如：sys_log/error/
func WriteLog(_filepath string, _txt string, _ip string, _filename string)  {
	// 创建文件夹
	filepath := helper.FrameInfo["storage_path"] + _filepath
	dateFile := helper.GetTimeDate("Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := helper.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	// 文件信息
	fileName := _filename + ".log"
	filePath := saveFilepath
	file, err := os.OpenFile(filePath+fileName, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// 延迟关闭文件
	defer file.Close()
	//
	date := helper.GetTimeDate("Y-m-d H:i:s")
	txt := "【" + date + "】" + _ip + "：" + _txt
	// 写入文件内容
	file.WriteString(txt + "\n")

	// 删除老文件夹
	_oldMonth := -32
	delFile := time.Now().AddDate(0, _oldMonth, -1).Format("20060102")
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("老文件夹删除失败=", err1)
	}
}