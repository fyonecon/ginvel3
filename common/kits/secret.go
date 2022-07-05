package kits
// 对称加密算法
// 调用：Encode(要加密的字符串, 混淆（可选）)、Decode(要解密的字符串, 混淆（可选）)
// 加密或解密失败，会返回空字符串

import (
	"ginvel/common/kits/kits_lib"
)


// Encode 加密
// 默认des，salt长度为8
func Encode(text string, salt string) string {
	sec := kits_lib.Secret{}
	if len(salt)<8 {salt += "fy2O22Il"}
	return sec.Encode(text, salt[0:8])
}
// Decode 解密
// 默认des，salt长度为8
func Decode(text string, salt string) string {
	sec := kits_lib.Secret{}
	if len(salt)<8 {salt += "fy2O22Il"}
	return sec.Decode(text, salt[0:8])
}


// EncodeDES 加密
func EncodeDES(text string, salt string) string {
	sec := kits_lib.SecretDES{}
	return string(sec.DesEncrypt([]byte(text), []byte(salt)))
}
// DecodeDES 解密
func DecodeDES(text string, salt string) string {
	sec := kits_lib.SecretDES{}
	return string(sec.DesDecrypt([]byte(text), []byte(salt)))
}


// EncodeAES 加密
func EncodeAES(text string, salt string) string {
	sec := kits_lib.SecretAES{}
	return string(sec.AesEncrypt([]byte(text), []byte(salt)))
}
// DecodeAES 解密
func DecodeAES(text string, salt string) string {
	sec := kits_lib.SecretAES{}
	return string(sec.AesDecrypt([]byte(text), []byte(salt)))
}
