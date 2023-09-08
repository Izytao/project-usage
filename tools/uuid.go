package tools

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// 生成Guid字串(32位)
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
