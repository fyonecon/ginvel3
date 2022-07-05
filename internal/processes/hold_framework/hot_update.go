package hold_framework

import (
	"fmt"
	"ginvel/internal/reader/http_block_json"
)

// HotUpdateRuler 热更服务
func (holdFramework *HoldFramework) HotUpdateRuler(num int)  {
	if num%8 == 0 { // 8个周期更新一次，30s*8=2400s。
		fmt.Println("运行动热更服务 >>> ", 8)
		http_block_json.ReadHttpBlockJson()
	}

	return
}
