/**
  @creator: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/10
  @desc: //TODO
**/

package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"strconv"
	"time"
)

var Key = "www.cilang.buzz1"

func AesEncrypt(plaintext []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	// PKCS#7 padding
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	// 创建CBC模式的、底层使用AES的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func GenerateTId(filename string) (string, error) {
	if filename == "" {
		filename = "cilang.buzz"
	}
	currentTime := time.Now()
	cipherText, err := AesEncrypt([]byte(strconv.Itoa(int(currentTime.Unix()))+filename), []byte(Key))
	if err != nil {
		log.Println("generate TId file", err)
		return "", err
	}
	encryptedData := base64.StdEncoding.EncodeToString(cipherText)

	return encryptedData, err
}

func GenerateToken() (string, error) {
	currentTime := time.Now().Add(48 * time.Hour)
	cipherText, err := AesEncrypt([]byte(strconv.Itoa(int(currentTime.Unix()))), []byte(Key))
	if err != nil {
		log.Println("generate token file", err)
		return "", err
	}
	encryptedData := base64.StdEncoding.EncodeToString(cipherText)
	return encryptedData, err
}
