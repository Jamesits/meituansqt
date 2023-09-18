package ecb

import (
	"crypto/aes"
)

func aesEncrypt(src []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(src) == 0 {
		return nil
	}
	ecb := NewECBEncrypter(block)
	dst := make([]byte, len(src))
	ecb.CryptBlocks(dst, src)
	return dst
}

func aesDecrypt(src []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(src) == 0 {
		return nil
	}
	ecb := NewECBDecrypter(block)
	dst := make([]byte, len(src))
	ecb.CryptBlocks(dst, src)
	return dst
}
