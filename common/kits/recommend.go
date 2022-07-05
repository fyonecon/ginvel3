package kits

import (
	"ginvel/common/kits/kits_lib"
)

// RecommendArraysId 分词ID推荐
// 多组id数组->
// 重复id组成优先级第一的新id数组1⃣️，不重复id各数组各取一个逐次组成新id数组2⃣️，目标id数组0⃣️=1⃣️+2⃣️
func RecommendArraysId(arrays map[string][]int) []int {
	rec := kits_lib.Recommend{}
	return rec.RecommendArraysInt(arrays)
}
