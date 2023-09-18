// Package sqtobfuscate implements the API request obfuscation required by Meituan / Dianping BEP API.
//
// The official documentation uses "signing" and "encryption" in a rather unclear way, and the algorithm described
// in the documentation is not signing nor encryption at all, thus we call it obfuscation here.
//
// Documentation: https://h5.dianping.com/app/bep-docs/sky-doc/api.html#_1-4-%E7%AD%BE%E5%90%8D%E6%96%B9%E6%B3%95
package sqtobfuscate

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"github.com/jamesits/meituansqt/pkg/ecb"
	"github.com/jamesits/meituansqt/pkg/pkcs5"
)

type Obfuscator struct {
	SecretKey string
}

func (o *Obfuscator) Obfuscate(srcString string) (string, error) {
	src := []byte(srcString)
	key, err := base64.StdEncoding.DecodeString(o.SecretKey)
	if err != nil {
		return "", fmt.Errorf("sqtobfuscate: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("sqtobfuscate: %w", err)
	}
	if len(src) == 0 {
		return "", nil
	}

	enc := ecb.NewECBEncrypter(block)
	src = pkcs5.Padding(src, enc.BlockSize())
	dst := make([]byte, len(src))
	enc.CryptBlocks(dst, src)

	return base64.RawURLEncoding.EncodeToString(dst), nil
}

func (o *Obfuscator) Deobfuscate(srcString string) (string, error) {
	src, err := base64.RawURLEncoding.DecodeString(srcString)
	if err != nil {
		return "", fmt.Errorf("sqtobfuscate: %w", err)
	}

	key, err := base64.StdEncoding.DecodeString(o.SecretKey)
	if err != nil {
		return "", fmt.Errorf("sqtobfuscate: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("sqtobfuscate: %w", err)
	}
	if len(src) == 0 {
		return "", nil
	}

	enc := ecb.NewECBDecrypter(block)
	dst := make([]byte, len(src))
	enc.CryptBlocks(dst, src)
	dst = pkcs5.Trimming(dst)
	return string(dst), err
}
