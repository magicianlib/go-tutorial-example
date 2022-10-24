package rsa_example

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Encrypt 公钥加密
// pubKeyFilename x509Der file
func Encrypt(pubDerFilename string, plaintext []byte) (ciphertext []byte, err error) {

	// 读取 x509 der 公钥文件数据
	der, err := os.ReadFile(pubDerFilename)
	if err != nil {
		return ciphertext, err
	}

	// PEM 解码, 得到 block 的 der 编码数据
	block, _ := pem.Decode(der)

	// 解码der得到公钥
	pkcs1, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return ciphertext, err
	}

	// 使用公钥进行数据加密
	ciphertext, err = rsa.EncryptPKCS1v15(rand.Reader, pkcs1, plaintext)
	return
}

// Decrypt 私钥解密
// privDerFilename x509Der file
func Decrypt(privDerFilename string, ciphertext []byte) (plaintext []byte, err error) {

	// 读取 x509 der 私钥文件数据
	der, err := os.ReadFile(privDerFilename)
	if err != nil {
		return plaintext, err
	}

	// PEM 解码, 得到 block 的 der 编码数据
	block, _ := pem.Decode(der)

	// 解码der得到私钥
	pkcs1, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return plaintext, err
	}

	// 使用私钥进行数据解密
	plaintext, err = rsa.DecryptPKCS1v15(rand.Reader, pkcs1, ciphertext)
	return

}
