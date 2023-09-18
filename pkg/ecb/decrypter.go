package ecb

import "crypto/cipher"

type ecbDecrypter struct{ cipher.Block }

// Deprecated: ECB is insecure, and you should not use it.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return ecbDecrypter{b}
}

func (x ecbDecrypter) BlockSize() int {
	return x.Block.BlockSize()
}

func (x ecbDecrypter) CryptBlocks(dst, src []byte) {
	size := x.BlockSize()
	if len(src)%size != 0 {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}
	if len(src) == 0 {
		return
	}

	for len(src) > 0 {
		x.Decrypt(dst, src)
		src, dst = src[size:], dst[size:]
	}
}
