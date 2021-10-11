package main

import (
	"crypto/rand"
	"fmt"
	"github.com/tjfoc/gmsm/x509"
	"log"
	"os"
)

func main() {
	//priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privPem, err := x509.WritePrivateKeyToPem(priv, nil)
	//fmt.Printf("%s\n", privPem)
	//pubKey, _ := priv.Public().(*sm2.PublicKey)
	//pubkeyPem, err := x509.WritePublicKeyToPem(pubKey)       // 生成公钥文件
	//fmt.Printf("%s\n", pubkeyPem)

	//fp, _ := os.Open("./test/test3/private.pem")
	//defer fp.Close()
	////测量文件长度以便于保存
	//fileinfo, _ := fp.Stat()
	//buf := make([]byte, fileinfo.Size())
	//fp.Read(buf)
	////下面的操作是与创建秘钥保存时相反的
	////pem解码
	////block, _ := pem.Decode(buf)
	//////x509解码,得到一个interface类型的pub
	////pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//priv, err := x509.ReadPrivateKeyFromPem(buf, nil) // 读取密钥
	//if err != nil {
	//	panic(err)
	//}
	//
	msg := []byte("QingTing321")
	//pub := &priv.PublicKey
	//ciphertxt, err := pub.EncryptAsn1(msg, rand.Reader) //sm2加密
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("加密结果:%x\n", ciphertxt)
	//fmt.Printf("加密结果:%s\n", base64.StdEncoding.EncodeToString(msg))
	//plaintxt, err := priv.DecryptAsn1(ciphertxt) //sm2解密
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if !bytes.Equal(msg, plaintxt) {
	//	log.Fatal("原文不匹配")
	//}
	//fmt.Printf("解密结果：%s\n", plaintxt)
	//
	//sign, err := priv.Sign(rand.Reader, msg, nil) //sm2签名
	//if err != nil {
	//	log.Fatal(err)
	//}
	//isok := pub.Verify(msg, sign) //sm2验签
	//fmt.Printf("Verified: %v\n", isok)

	decryptData := encrypt(msg)
	fmt.Println(111)
	fmt.Printf("%s\n", decryptData)
	fmt.Println(111)

	data := decrypt(decryptData)
	fmt.Println(111)
	fmt.Printf("%s\n", data)
}

func decrypt(msg []byte) []byte {
	fp, _ := os.Open("./test/test3/private.pem")
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	//block, _ := pem.Decode(buf)
	////x509解码,得到一个interface类型的pub
	//pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	priv, err := x509.ReadPrivateKeyFromPem(buf, nil) // 读取密钥
	if err != nil {
		panic(err)
	}

	res, err := priv.DecryptAsn1(msg)
	if err != nil {
		log.Fatal(err)
	}

	//pubKey, _ := priv.Public().(*sm2.PublicKey)
	//
	//ciphertxt, err := pubKey.EncryptAsn1(msg, rand.Reader) //sm2加密
	//if err != nil {
	//	log.Fatal(err)
	//}

	return res
}

func encrypt(msg []byte) []byte {
	fp, _ := os.Open("./test/test3/public.pem")
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	//block, _ := pem.Decode(buf)
	////x509解码,得到一个interface类型的pub
	//pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	fmt.Printf("%v", string(buf))
	fmt.Println(12213213213213)
	pub, err := x509.ReadPublicKeyFromPem(buf) // 读取密钥
	if err != nil {
		panic(err)
	}

	ciphertxt, err := pub.EncryptAsn1(msg, rand.Reader) //sm2加密
	if err != nil {
		log.Fatal(err)
	}

	return ciphertxt
}
