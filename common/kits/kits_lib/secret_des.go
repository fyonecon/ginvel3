package kits_lib

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

type SecretDES struct {}

//dec的CBC加密
//编写填充函数，如果最后一个分组字节数不够，填充
//字节数刚好合适，添加一个新的分组
//填充个的字节的值==缺少的字节的数
func (secret *SecretDES) PaddingLastGroup(plainText []byte, blockSize int) []byte {
	//1. 求出最后一个组中需要填充的字节数  每组bloclSize字节
	padNum := blockSize - len(plainText) % 8
	//2. 创建切片，长度=padNum   每个字节值byte(padNum)
	char := []byte{byte(padNum)}//长度1
	//切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	//3. newPlain数组追加到原始明文的后边
	newPlain = append(plainText, newPlain...)
	return newPlain
}

func (secret *SecretDES) unPaddingLastGroup(plaintext []byte) []byte{
	//1. 取出切片最后一个字节内容
	length := len(plaintext)
	lastChar := plaintext[length-1] //byte类型
	number := int(lastChar)  //获取尾部填充字节个数
	return plaintext[:length-number]
}

//DesEncrypt des加密 分组模式cbc
func (secret *SecretDES) DesEncrypt(plaintext, key []byte) []byte {
	//1. 建立一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 明文填充
	newText := secret.PaddingLastGroup(plaintext, block.BlockSize())
	//3. 创建一个使用cbc分组的接口
	iv := []byte("12345678")//初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4. 加密 将newText加密存储到cipherText
	cipherText := make([]byte,len(newText))
	blockMode.CryptBlocks(cipherText,newText)
	return cipherText
}

//DesDecrypt des解密
func (secret *SecretDES) DesDecrypt(cipherText, key []byte) []byte  {
	//1. 建立一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2. 创建一个使用cbc分组的接口
	iv := []byte("12345678")//初始化向量
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3. 解密
	newText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(newText, cipherText)

	//4. 去掉尾部
	plainText := secret.unPaddingLastGroup(newText)
	return plainText
}
