package ecb

// code from: https://stackoverflow.com/a/66512550

import (
	"crypto/cipher"
)

type ecbEncrypter struct{ cipher.Block }

// Deprecated: ECB is insecure, and you should not use it.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return ecbEncrypter{b}
}

func (x ecbEncrypter) BlockSize() int {
	return x.Block.BlockSize()
}

func (x ecbEncrypter) CryptBlocks(dst, src []byte) {
	size := x.BlockSize()
	if len(src)%size != 0 {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}

	for len(src) > 0 {
		x.Encrypt(dst, src)
		src, dst = src[size:], dst[size:]
	}
}
