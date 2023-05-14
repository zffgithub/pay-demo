package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/*
* 获取小写的MD5
 */
func GetMD5LOWER(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/*
* 获取大写的MD5
 */
func GetMD5Upper(s string) string {
	return strings.ToUpper(GetMD5LOWER(s))
}
