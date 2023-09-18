package ecb

import (
	"crypto/aes"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testSet struct {
	key    string
	plain  string
	cipher string
}

// https://github.com/ircmaxell/quality-checker/blob/63e91ea991ad98428e2bcced2c35492ca634cc28/tmp/gh_18/PHP-PasswordLib-master/test/Data/Vectors/aes-ecb.test-vectors
var aesTestSets = []testSet{
	{
		key:    "2b7e151628aed2a6abf7158809cf4f3c",
		plain:  "6bc1bee22e409f96e93d7e117393172a",
		cipher: "3ad77bb40d7a3660a89ecaf32466ef97",
	},
	{
		key:    "2b7e151628aed2a6abf7158809cf4f3c",
		plain:  "ae2d8a571e03ac9c9eb76fac45af8e51",
		cipher: "f5d3d58503b9699de785895a96fdbaaf",
	},

	{
		key:    "2b7e151628aed2a6abf7158809cf4f3c",
		plain:  "30c81c46a35ce411e5fbc1191a0a52ef",
		cipher: "43b1cd7f598ece23881b00e3ed030688",
	},
	{
		key:    "2b7e151628aed2a6abf7158809cf4f3c",
		plain:  "f69f2445df4f9b17ad2b417be66c3710",
		cipher: "7b0c785e27e8ad3f8223207104725dd4",
	},
	{
		key:    "8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b",
		plain:  "6bc1bee22e409f96e93d7e117393172a",
		cipher: "bd334f1d6e45f25ff712a214571fa5cc",
	},
	{
		key:    "8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b",
		plain:  "ae2d8a571e03ac9c9eb76fac45af8e51",
		cipher: "974104846d0ad3ad7734ecb3ecee4eef",
	},
	{
		key:    "8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b",
		plain:  "30c81c46a35ce411e5fbc1191a0a52ef",
		cipher: "ef7afd2270e2e60adce0ba2face6444e",
	},
	{
		key:    "8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b",
		plain:  "f69f2445df4f9b17ad2b417be66c3710",
		cipher: "9a4b41ba738d6c72fb16691603c18e0e",
	},
	{
		key:    "603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		plain:  "6bc1bee22e409f96e93d7e117393172a",
		cipher: "f3eed1bdb5d2a03c064b5a7e3db181f8",
	},
	{
		key:    "603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		plain:  "ae2d8a571e03ac9c9eb76fac45af8e51",
		cipher: "591ccb10d410ed26dc5ba74a31362870",
	},
	{
		key:    "603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		plain:  "30c81c46a35ce411e5fbc1191a0a52ef",
		cipher: "b6ed21b99ca6f4f9f153e7b1beafed1d",
	},
	{
		key:    "603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		plain:  "f69f2445df4f9b17ad2b417be66c3710",
		cipher: "23304b7a39f9f3ff067d8d8f9e24ecc7",
	},
}

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

func TestAESEncryption(t *testing.T) {
	for _, set := range aesTestSets {
		key, err := hex.DecodeString(set.key)
		assert.NoError(t, err)

		plain, err := hex.DecodeString(set.plain)
		assert.NoError(t, err)

		cipher, err := hex.DecodeString(set.cipher)
		assert.NoError(t, err)

		enc := aesEncrypt(plain, key)
		assert.EqualValues(t, cipher, enc)
	}
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

func TestAESDecryption(t *testing.T) {
	for _, set := range aesTestSets {
		key, err := hex.DecodeString(set.key)
		assert.NoError(t, err)

		plain, err := hex.DecodeString(set.plain)
		assert.NoError(t, err)

		cipher, err := hex.DecodeString(set.cipher)
		assert.NoError(t, err)

		dec := aesDecrypt(cipher, key)
		assert.EqualValues(t, plain, dec)
	}
}
