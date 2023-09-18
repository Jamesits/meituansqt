package sqtobfuscate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObfuscate(t *testing.T) {
	for _, set := range sqtTestSets {
		ob := Obfuscator{SecretKey: set.key}

		enc, err := ob.Obfuscate(set.plain)
		assert.NoError(t, err)
		assert.EqualValues(t, set.cipher, enc)
	}
}

func TestDeobfuscate(t *testing.T) {
	for _, set := range sqtTestSets {
		ob := Obfuscator{SecretKey: set.key}

		dec, err := ob.Deobfuscate(set.cipher)
		assert.NoError(t, err)
		assert.EqualValues(t, set.plain, dec)
	}
}
