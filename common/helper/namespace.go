package helper

import (
	"ginvel/internal/reader/framework_toml"
)

// 空间命名下的全局变量
var frameworkConfig = framework_toml.ReadFrameworkToml()
