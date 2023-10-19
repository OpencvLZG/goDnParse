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
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES解密
func AesDecrypt(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)

	// 去除PKCS#7填充
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// 去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return origData
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DeEncryptTId(tid string) ([]byte, error) {
	encrypted, err := base64.StdEncoding.DecodeString(tid)
	if err != nil {
		return []byte(""), err
	}
	decrypted, err := AesDecrypt(encrypted, []byte(Key))
	if err != nil {
		return []byte(""), err
	}
	return decrypted, err
}
