package ecb

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
