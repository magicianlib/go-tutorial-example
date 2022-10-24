package rsa_example

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func HashWithRsaSign(privDerFilename string, hash crypto.Hash, data []byte) (base64sign string, err error) {

	// 读取 x509 der 私钥文件数据
	der, err := os.ReadFile(privDerFilename)
	if err != nil {
		return "", err
	}

	// PEM 解码, 得到 block 的 der 编码数据
	block, _ := pem.Decode(der)
	pkcs1, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Hash 计算, 得到签名
	h := hash.New()
	h.Write(data)
	h.Sum(nil)

	b, err := rsa.SignPKCS1v15(rand.Reader, pkcs1, hash, h.Sum(nil))
	if err != nil {
		return "", err
	}

	// 转 base64
	return base64.StdEncoding.EncodeToString(b), nil
}

func HashWithRsaSignVerify(pubDerFilename string, hash crypto.Hash, data []byte, base64sign string) error {

	// 读取 x509 der 公钥文件数据
	der, err := os.ReadFile(pubDerFilename)
	if err != nil {
		return err
	}

	// PEM 解码, 得到 block 的 der 编码数据
	block, _ := pem.Decode(der)
	pkcs1, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// Hash 计算, 得到签名
	h := hash.New()
	h.Write(data)
	h.Sum(nil)

	sign, err := base64.StdEncoding.DecodeString(base64sign)
	if err != nil {
		return err
	}

	// 验证签名
	err = rsa.VerifyPKCS1v15(pkcs1, hash, h.Sum(nil), sign)
	if err != nil {
		return err
	}

	return nil
}
