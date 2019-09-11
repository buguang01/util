package util

import (
	"encoding/hex"
	"crypto/md5"
)

//生成MD5码
func GetMd5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}