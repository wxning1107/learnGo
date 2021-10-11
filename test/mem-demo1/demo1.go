package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

// des的CBC加密
// 编写填充函数, 如果最后一个分组字节数不够, 填充
// ......字节数刚好合适, 添加一个新的分组
// 填充个的字节的值 == 缺少的字节的数
func paddingLastGroup(plainText []byte, bloclSize int) []byte {
	// 1. 求出最后一个组中剩余的字节数 28 % 8 = 3...4  32 % 8 = 4 ...0
	padNum := bloclSize - len(plainText)%bloclSize
	// 2. 创建新的切片, 长度 == padNum, 每个字节值 byte(padNum)
	char := []byte{byte(padNum)} // 长度1,
	// 切片创建, 并初始化
	newPlain := bytes.Repeat(char, padNum)
	// 3. newPlain数组追加到原始明文的后边
	newText := append(plainText, newPlain...)
	return newText
}

// 去掉填充的数据
func unPaddingLastGrooup(plainText []byte) []byte {
	// 1. 拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1] //
	number := int(lastChar)         // 尾部填充的字节个数
	return plainText[:length-number]
}

// des加密
func desEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 3. 创建一个使用cbc分组接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4. 加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)
	// blockMode.CryptBlocks(newText, newText)
	return cipherText
}

// des解密
func desDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用cbc模式解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 4. cipherText现在存储的是明文, 需要删除加密时候填充的尾部数据
	plainText := unPaddingLastGrooup(cipherText)
	return plainText
}

// aes加密, 分组模式ctr
func aesEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr分组接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)

	// 4. 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return []byte(base64.StdEncoding.EncodeToString(cipherText))
}

// des解密
func aesDecrypt(cipherText, key []byte) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(string(cipherText))
	// 1. 建一个底层使用des的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr模式解密的接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 3. 解密
	stream.XORKeyStream(decodeString, decodeString)

	return decodeString
}

func main() {
	//fmt.Println("des 加解密")
	//key := []byte("1234abdd")
	src := []byte("QingTing321")
	//cipherText := desEncrypt(src, key)
	//fmt.Printf("加密后的数据: %v\n", string(cipherText))
	//plainText := desDecrypt(cipherText, key)
	//fmt.Printf("解密之后的数据: %s\n", string(plainText))

	key1 := []byte("1234abdd12345678")
	cipherText := aesEncrypt(src, key1)
	fmt.Printf("加密后的数据: %s\n", string(cipherText))
	plainText := aesDecrypt(cipherText, key1)
	fmt.Printf("解密之后的数据: %s\n", string(plainText))
}
