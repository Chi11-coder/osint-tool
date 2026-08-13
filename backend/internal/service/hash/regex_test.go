package hash_test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"example.com/security/internal/service/hash"
)

func TestFileHash(t *testing.T) {
	hashBytes := sha256.Sum256([]byte("Test"))
	getHash := hex.EncodeToString(hashBytes[:])
	datas := []struct {
		title    string
		testHash string
		want     bool
	}{
		{
			"正常系: sha256ハッシュ",
			getHash,
			true,
		},
		{
			"異常系: 64文字未満",
			getHash[0:63],
			false,
		},
		{
			"異常系: 65文字以上",
			getHash + "1",
			false,
		},
		{
			"異常系: 不正な文字列",
			getHash[0:63] + "z",
			false,
		},
	}

	for _, data := range datas {

		if data.want {
			t.Run(data.title, func(t *testing.T) {
				if ok := hash.IsValid(data.testHash); ok != true {
					t.Errorf("expected true. but %v is failed regex", data.testHash)
				}
			})
		} else {
			t.Run(data.title, func(t *testing.T) {
				if ok := hash.IsValid(data.testHash); ok != false {
					t.Errorf("expected false. but %v is true regex", data.testHash)
				}
			})
		}
	}
}
