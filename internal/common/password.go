package common

import (
	"github.com/donech/tool/cipher"
)

var key string = "1234567890123456"

func SetKey(k string) {
	if key == "1234567890123456" && k != "" {
		key = k
	}
}

func SignPassword(password string) string {
	res := cipher.AesEncryptCBC([]byte(password), []byte(key))
	return string(res)
}
