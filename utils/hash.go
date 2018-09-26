package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// cbc 解密
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if len(ciphertext) < aes.BlockSize {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
