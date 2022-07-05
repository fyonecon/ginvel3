package helper

import (
	"runtime"
)

// FrameInfo 框架基本参数，承接参数
var FrameInfo = map[string]string{
	"framework_hash": MD5("22O2levinG") + "+hash", // 与toml文件密钥对应一致，自定义。这将编译在代码里，如果与toml文件中的不匹配，将中断运行。
	"timezone": frameworkConfig.Framework.Timezone, // 时区
	"gv_version": frameworkConfig.Framework.GlVersion,
	"go_version": runtime.Version(),
	"go_root": runtime.GOROOT(),
	"framework_path": frameworkConfig.Framework.FrameworkPath,
	"storage_path": frameworkConfig.Framework.FrameworkPath + "storage/",
	"env": frameworkConfig.Common.Env,
	"http_host": frameworkConfig.HttpServer.Host,
	"http_port": frameworkConfig.HttpServer.Port,
}

// Paging 分页参数
var Paging = map[string]int64{
	"limit": 20, // 每页多少条数据
	"max_page": 1000, // 最大查到多少页
	"min_limit": 2, // 最少一次查多少条数据
	"max_limit": 100, // 最多一次查多少条数据
}

// AdminPwd 管理员登录密码
func AdminPwd(pwd string) string {
	return MD5(SaltPwdAdmin(pwd+"Adfy0ne"))
}

// UserPwd 用户登录密码
func UserPwd(pwd string) string {
	return MD5(SaltPwdUser(pwd+"UsfyOne"))
}
