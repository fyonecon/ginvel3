package http_block_json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var BlackIp []string
var WhiteIp []string
var MaxSniffHourIpPV int64
var MaxDDOSHourIpPV int64

// ReadHttpBlockJson 读json文件，并赋值全局变量
func ReadHttpBlockJson() HttpBlockConfig {
	// main.go文件的绝对路径
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"

	// json文件及其数据模型
	var file = mainDirectory + "storage/config_json/http_block.json"
	var httpBlockConfig HttpBlockConfig

	/*

	setescapehtml=false

	bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(js)

	*/

	// 方法1
	//var fileLocker sync.Mutex
	//fileLocker.Lock()
	//dataJson, _ := ioutil.ReadFile(file)
	//fileLocker.Unlock()
	//err := json.Unmarshal([]byte(dataJson), &httpBlockConfig)

	// 方法2
	dataJson, _ := os.Open(file)
	defer dataJson.Close()
	decoder := json.NewDecoder(dataJson)
	err := decoder.Decode(&httpBlockConfig)

	if err != nil {
		log.Println("'http_block.json' ：Decoder Error = ", err.Error(), "httpBlockJson参数未拿到（跳过，将使用默认值）")
		//os.Exit(200)
	}

	_blackIp := httpBlockConfig.BlackIp
	_whiteIp := httpBlockConfig.WhiteIp

	BlackIp = strings.Split(_blackIp, "#@")
	WhiteIp = strings.Split(_whiteIp, "#@")
	MaxSniffHourIpPV = httpBlockConfig.MaxSniffHourIpPV
	MaxDDOSHourIpPV = httpBlockConfig.MaxDDOSHourIpPV

	fmt.Println("ReadHttpBlockJson=", BlackIp, WhiteIp)

	return httpBlockConfig
}