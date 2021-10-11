package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func Getkeys() {
	//得到私钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_Privatekey := x509.MarshalPKCS1PrivateKey(privateKey)
	//创建一个用来保存私钥的以.pem结尾的文件
	fp, _ := os.Create("./test/test2/RSA_private.pem")
	defer fp.Close()
	//将私钥字符串设置到pem格式块中
	pem_block := pem.Block{
		Type:  "RSA_privateKey",
		Bytes: x509_Privatekey,
	}
	//转码为pem并输出到文件中
	pem.Encode(fp, &pem_block)

	//处理公钥,公钥包含在私钥中
	publickKey := privateKey.PublicKey
	//接下来的处理方法同私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_PublicKey, _ := x509.MarshalPKIXPublicKey(&publickKey)
	pem_PublickKey := pem.Block{
		Type:  "RSA_PublicKey",
		Bytes: x509_PublicKey,
	}
	file, _ := os.Create("./test/test2/RSA_PublicKey.pem")
	defer file.Close()
	//转码为pem并输出到文件中
	pem.Encode(file, &pem_PublickKey)

}

//使用公钥进行加密
func RSA_encrypter(path string, msg []byte) []byte {
	//首先从文件中提取公钥
	fp, _ := os.Open(path)
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	fmt.Printf("%s\n", buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码,得到一个interface类型的pub
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型
	cipherText, _ := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), msg)

	return []byte(base64.StdEncoding.EncodeToString(cipherText))
}

//使用私钥进行解密
func RSA_decrypter(path string, cipherText []byte) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(string(cipherText))
	//同加密时，先将私钥从文件中取出，进行二次解码
	fp, _ := os.Open(path)
	defer fp.Close()
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	block, _ := pem.Decode(buf)
	PrivateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//二次解码完毕，调用解密函数
	afterDecrypter, _ := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, decodeString)

	return afterDecrypter
}

func main() {
	//尝试调用
	msg := []byte("QingTing321")
	ciphertext := RSA_encrypter("./test/test2/RSA_PublicKey.pem", msg)
	//转化为十六进制方便查看结果
	fmt.Printf("加密结果：%s\n", string(ciphertext))
	result := RSA_decrypter("./test/test2/RSA_private.pem", ciphertext)
	fmt.Printf("解密结果：%s\n", string(result))
	//Getkeys()
}
