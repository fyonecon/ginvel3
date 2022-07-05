package middlewares

import (
	"fmt"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"ginvel/internal/reader/http_block_json"
	"github.com/gin-gonic/gin"
	"net"
	"net/url"

	//"net/url"
	"strings"
)

// 针对DDOS和恶意嗅探的拦截
// 返回404、永久拉黑IP
// IP-PV（白名单第三方接口访问除外）：每日>=4990限制24h[暂关]，每日>=9990永封；正则黑名单词仅GET：每日>=190限制24h，累计490永封[暂关]；访问频率不限制；访问接口及方法不限制。

// StringArrayValueInString 数组中某个值是否在string中
func StringArrayValueInString(bigString string, stringArray []string) bool {
	bigString = strings.ToLower(bigString) // 转成小写
	for i:=0; i<len(stringArray); i++ {
		theValue := stringArray[i]
		if strings.Index(bigString, theValue) != -1{ // 包含
			return true
		}
	}
	return false
}

// StringArrayHasValue string数组是否包含某值
func StringArrayHasValue(stringArray []string, value string) int64 {
	for i:=0; i<len(stringArray); i++ {
		if stringArray[i] == value {
			return int64(i)
		}
	}
	return -1
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

// BlockSniff 匹配正则词，仅限GET方法的url
// false正常访问，true被拦截
func BlockSniff(ctx *gin.Context, uri string, ip string) bool {
	var maxHourIpPV int64 = http_block_json.MaxSniffHourIpPV // pv
	if maxHourIpPV <= 0 { maxHourIpPV = 190 }
	var blackWord []string = []string{ // 此值不需要热更：1。部分字符转义与fresh相斥，过滤的字段是固定的；2。请写成小写。
		"myadmin", "exit", "function", "_path", "select", "table",
		"include", "xss", "etc", ".ini",
		"%20and%20", "nslookup", "%%", "%00", "%27", "union", "eval", "response", "action",
		"echo", "invoke", "md5(", "driver", "ashx", "))", "passwd",  "?*", "**", "()", "shell", "char", "concat", "white",
		"hex", "uploads", "sqladmin", ".aspx", ".asp", "wordpress", "think", "cache", "manifest", ".top",
	}

	hasSniff := StringArrayValueInString(uri, blackWord)
	if hasSniff == true { // 包含
		// 记录访问IP
		var updateIpPV, _ = kits.UpdateIP(ctx, "sniff-" + ip)
		fmt.Println("IP=", ip, "，PV=sniff=", updateIpPV)
		return true
	}

	// 查询IP-pv
	var allIpPV, hourIpPV = kits.SearchIP(ctx, "sniff-" + ip)
	fmt.Println("查询IP-PV=search-sniff=", ip, allIpPV, hourIpPV)

	if hourIpPV > maxHourIpPV || hourIpPV < 0{
		return true
	}else {
		return false
	}
}

// BlockDDOS 拦截DDOS
// false正常访问，true被拦截
// 白名单IP除外
func BlockDDOS(ctx *gin.Context, ip string) bool {
	var maxHourIpPV int64 = http_block_json.MaxDDOSHourIpPV // pv
	if maxHourIpPV <= 0 { maxHourIpPV = 4990 }

	// 黑名单IP
	var blackIp []string = http_block_json.BlackIp
	if len(blackIp) == 0{
		blackIp = []string{
			"192.168.0.0", "192.168.0.1",
		}
	}

	// 排除白名单IP
	var whiteIp []string = http_block_json.WhiteIp
	if len(whiteIp) == 0{
		whiteIp = []string{
			"::0", "0.0.0.0", "127.0.0.1",
		}
	}

	// 服务器IP
	var serverIP string = ServerIPv4()[0]

	var hasBlackValue int64 = StringArrayHasValue(blackIp, ip)
	var hasWhiteValue int64 = StringArrayHasValue(whiteIp, ip)
	if hasBlackValue != -1 { // IP在黑名单
		return true
	}else {
		if hasWhiteValue != -1 || ip == serverIP { // 在白名单
			// 记录访问IP
			var updateIpPV, _ = kits.UpdateIP(ctx, "ddos-" + ip)
			fmt.Println("IP=", ip, "，PV=ddos=", updateIpPV)
			return false
		}

		// 查询IP-pv
		var allIpPV, hourIpPV = kits.SearchIP(ctx, "ddos-" + ip)
		fmt.Println("查询IP-PV=search=ddos=", ip, allIpPV, hourIpPV)

		if hourIpPV > maxHourIpPV || hourIpPV < 0{
			return true
		}else {
			return false
		}
	}
}

// CheckBlockSniff level-1
func CheckBlockSniff(ctx *gin.Context, uri string, ip string)  {
	// Sniff
	if BlockSniff(ctx, uri, ip) == false { // 正常访问
		ctx.Next()
	}else {
		// 返回
		uri = url.QueryEscape(uri)
		ctx.JSON(404, gin.H{
			"state": 404,
			"msg": "触及规则限制",
			"content": gin.H{
				"gv_version": helper.FrameInfo["gv_version"],
				"go_version": helper.FrameInfo["go_version"],
				"uri": uri,
			},
		})
		ctx.Abort()
	}

	return
}

// CheckBlockDDOS level-2
func CheckBlockDDOS(ctx *gin.Context, ip string)  {
	// DDOS
	if BlockDDOS(ctx, ip) == false { // 正常访问
		ctx.Next()
	}else {
		// 返回
		ctx.JSON(404, gin.H{
			"state": 404,
			"msg": "触及屏蔽",
			"content": gin.H{
				"gv_version": helper.FrameInfo["gv_version"],
				"go_version": helper.FrameInfo["go_version"],
				"ip": ip,
			},
		})
		ctx.Abort()
	}

	return
}

// HttpBlock 拦截http非法恶意访问
func HttpBlock(ctx *gin.Context) {
	var uri string = "" // 用户访问网址
	var ip string = "" // 用户访问IP

	_host, _ := ctx.Get("host")
	host := helper.ValueInterfaceToString(_host)
	uri = host + ctx.Request.RequestURI

	ip = ctx.ClientIP()

	CheckBlockSniff(ctx, uri, ip)
	CheckBlockDDOS(ctx, ip)

	return
}
