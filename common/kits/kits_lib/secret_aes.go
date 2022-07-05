package kits_lib

import (
	"crypto/aes"
	"crypto/cipher"
)

type SecretAES struct {}

// AesEncrypt aes加密 分组模式ctr
func (secret *SecretAES) AesEncrypt(plaintext, key []byte) []byte {
	//1. 建立一个底层使用aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 明文填充
	//newText := PaddingLastGroup(plaintext, block.BlockSize())
	//3. 创建一个使用ctr分组的接口
	iv := []byte("12345678abcdefgh") // 盐
	stream := cipher.NewCTR(block, iv)
	//4. 加密 将plaintext加密存储到cipherText
	cipherText := make([]byte,len(plaintext))
	stream.XORKeyStream(cipherText,plaintext)
	return cipherText
}

// AesDecrypt aes解密
func (secret *SecretAES) AesDecrypt(cipherText, key []byte) []byte  {
	//1. 建立一个底层使用des的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2. 创建一个使用cbc分组的接口
	iv := []byte("12345678abcdefgh") // 盐
	stream := cipher.NewCTR(block, iv)
	//3. 解密
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText,cipherText)

	return plainText
}
